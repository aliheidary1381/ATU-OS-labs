// go test with "-race" flag
package kvserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gocheck "gopkg.in/check.v1"
	"server/internal/pb"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	gocheck.TestingT(t)
}

type TestSuite struct {
	connection *grpc.ClientConn
	db         pb.DBClient
	ctx        context.Context
}

var _ = gocheck.Suite(&TestSuite{})

func (t *TestSuite) SetUpSuite(c *gocheck.C) {
	var err error
	t.ctx = context.Background()
	serverAddress := fmt.Sprintf("localhost:%d", *Port)
	t.connection, err = grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	c.Assert(err, gocheck.IsNil, gocheck.Commentf("test did not connect to server: %v", err))
	t.db = pb.NewDBClient(t.connection)

	res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "0"})
	if err != nil {
		if len(err.Error()) > 148 && err.Error()[130:148] == "connection refused" {
			c.Fatalf("test did not connect to server: %v", err)
		}
	}
	c.Assert(err, gocheck.IsNil)
	c.Assert(res.StatusCode, gocheck.Equals, 200)
}

func (t *TestSuite) TestSet(c *gocheck.C) {
	res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "0"})
	c.Assert(err, gocheck.IsNil)
	c.Assert(res.StatusCode, gocheck.Equals, 200)
}

func (t *TestSuite) TestWeirdSet(c *gocheck.C) {
	res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "'1%*x", Value: "🙂"})
	c.Assert(err, gocheck.IsNil)
	c.Assert(res.StatusCode, gocheck.Equals, 200)
}

func (t *TestSuite) TestSetGet(c *gocheck.C) {
	res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "0"})
	c.Assert(err, gocheck.IsNil)
	c.Assert(res.StatusCode, gocheck.Equals, 200)
	res2, err2 := t.db.Get(t.ctx, &pb.GetRequest{Key: "x"})
	c.Assert(err2, gocheck.IsNil)
	c.Assert(res2.Value, gocheck.Equals, "0")
}

func (t *TestSuite) TestWeirdSetGet(c *gocheck.C) {
	res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "'1%*x", Value: "🙂"})
	c.Assert(err, gocheck.IsNil)
	c.Assert(res.StatusCode, gocheck.Equals, 200)
	res2, err2 := t.db.Get(t.ctx, &pb.GetRequest{Key: "'1%*x"})
	c.Assert(err2, gocheck.IsNil)
	c.Assert(res2.Value, gocheck.Equals, "🙂")
}

func (t *TestSuite) TestSetSetGet(c *gocheck.C) {
	for i := 0; i < 5; i++ {
		res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "0"})
		c.Assert(err, gocheck.IsNil)
		c.Assert(res.StatusCode, gocheck.Equals, 200)
		res, err = t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "1"})
		c.Assert(err, gocheck.IsNil)
		c.Assert(res.StatusCode, gocheck.Equals, 200)
		res2, err2 := t.db.Get(t.ctx, &pb.GetRequest{Key: "x"})
		c.Assert(err2, gocheck.IsNil)
		c.Assert(res2.Value, gocheck.Equals, "1")
	}
}

func (t *TestSuite) TestGetNotSet(c *gocheck.C) {
	res, err := t.db.Get(t.ctx, &pb.GetRequest{Key: "x"})
	c.Assert(err, gocheck.IsNil)
	c.Assert(res.Value, gocheck.Equals, "")
}

func (t *TestSuite) TestConcurrentSetOnOneKey(c *gocheck.C) {
	wg := sync.WaitGroup{}
	wg.Add(200)
	defer wg.Wait()
	for i := 0; i < 200; i++ {
		go func() {
			res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "0"})
			c.Assert(err, gocheck.IsNil)
			c.Assert(res.StatusCode, gocheck.Equals, 200)
		}()
	}
}

func (t *TestSuite) TestSetConcurrentGetOnOneKey(c *gocheck.C) {
	wg := sync.WaitGroup{}
	wg.Add(200)
	defer wg.Wait()
	res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "0"})
	c.Assert(err, gocheck.IsNil)
	c.Assert(res.StatusCode, gocheck.Equals, 200)
	for i := 0; i < 200; i++ {
		go func() {
			res, err := t.db.Get(t.ctx, &pb.GetRequest{Key: "x"})
			c.Assert(err, gocheck.IsNil)
			c.Assert(res.Value, gocheck.Equals, "0")
		}()
	}
}

func (t *TestSuite) TestConcurrentSet(c *gocheck.C) {
	wg := sync.WaitGroup{}
	wg.Add(200)
	defer wg.Wait()
	for _, ch := range [5]string{"a", "b", "c", "d", "e"} {
		go func(ch string) {
			for i := 0; i < 40; i++ {
				go func() {
					res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: ch, Value: "0"})
					c.Assert(err, gocheck.IsNil)
					c.Assert(res.StatusCode, gocheck.Equals, 200)
				}()
			}
		}(ch)
	}
}

func (t *TestSuite) TestSetConcurrentGet(c *gocheck.C) {
	wg := sync.WaitGroup{}
	wg.Add(200)
	defer wg.Wait()
	for _, ch := range [5]string{"a", "b", "c", "d", "e"} {
		go func(ch string) {
			res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: ch, Value: "0"})
			c.Assert(err, gocheck.IsNil)
			c.Assert(res.StatusCode, gocheck.Equals, 200)
			for i := 0; i < 40; i++ {
				go func() {
					res, err := t.db.Get(t.ctx, &pb.GetRequest{Key: ch})
					c.Assert(err, gocheck.IsNil)
					c.Assert(res.Value, gocheck.Equals, "0")
				}()
			}
		}(ch)
	}
}

func (t *TestSuite) TestConcurrentSetSetGet(c *gocheck.C) {
	wg := sync.WaitGroup{}
	wg.Add(200)
	defer wg.Wait()
	for _, ch := range [5]string{"a", "b", "c", "d", "e"} {
		go func(ch string) {
			for i := 0; i < 40; i++ {
				res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: ch, Value: "0"})
				c.Assert(err, gocheck.IsNil)
				c.Assert(res.StatusCode, gocheck.Equals, 200)
				res, err = t.db.Set(t.ctx, &pb.SetRequest{Key: ch, Value: "1"})
				c.Assert(err, gocheck.IsNil)
				c.Assert(res.StatusCode, gocheck.Equals, 200)
				res2, err2 := t.db.Get(t.ctx, &pb.GetRequest{Key: ch})
				c.Assert(err2, gocheck.IsNil)
				c.Assert(res2.Value, gocheck.Equals, "1")
			}
		}(ch)
	}
}

func (t *TestSuite) TestConcurrentSetAndGet(c *gocheck.C) {
	wg := sync.WaitGroup{}
	wg.Add(200)
	defer wg.Wait()
	res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "0"})
	c.Assert(err, gocheck.IsNil)
	c.Assert(res.StatusCode, gocheck.Equals, 200)
	for i := 0; i < 100; i++ {
		go func() {
			res, err := t.db.Set(t.ctx, &pb.SetRequest{Key: "x", Value: "0"})
			c.Assert(err, gocheck.IsNil)
			c.Assert(res.StatusCode, gocheck.Equals, 200)
		}()
		go func() {
			res, err := t.db.Get(t.ctx, &pb.GetRequest{Key: "x"})
			c.Assert(err, gocheck.IsNil)
			c.Assert(res.Value, gocheck.Equals, "0")
		}()
	}
}

func (t *TestSuite) TearDownSuite(_ *gocheck.C) {
	_ = t.connection.Close()
}
