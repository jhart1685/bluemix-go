package k8sclusterv1

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Bluemix/bluemix-go"
	"github.com/IBM-Bluemix/bluemix-go/client"
	bluemixHttp "github.com/IBM-Bluemix/bluemix-go/http"
	"github.com/IBM-Bluemix/bluemix-go/session"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Workers", func() {
	var server *ghttp.Server
	Describe("Add", func() {
		Context("When adding a worker is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/workers"),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should return worker added to cluster", func() {
				target := &ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WorkerParam{
					Action: "add", Count: 1,
				}
				err := newWorker(server.URL()).Add("test", params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When adding worker is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/workers"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to add worker to cluster`),
					),
				)
			})

			It("should return error during add webhook to cluster", func() {
				target := &ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WorkerParam{
					Action: "add", Count: 1,
				}
				err := newWorker(server.URL()).Add("test", params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Get
	Describe("Get", func() {
		Context("When retrieving available workers is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusOK, `{"Billing":"","ErrorMessage":"","Isolation":"","MachineType":"free","KubeVersion":"","PrivateIP":"","PublicIP":"","PrivateVlan":"vlan","PublicVlan":"vlan","state":"normal","status":"ready"}`),
					),
				)
			})

			It("should return available workers ", func() {
				target := &ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				worker, err := newWorker(server.URL()).Get("abc-123-def-ghi", target)
				Expect(err).NotTo(HaveOccurred())
				Expect(worker).ShouldNot(BeNil())
			})
		})
		Context("When retrieving available workers is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve workers`),
					),
				)
			})

			It("should return error during retrieveing workers", func() {
				target := &ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				worker, err := newWorker(server.URL()).Get("abc-123-def-ghi", target)
				Expect(err).To(HaveOccurred())
				Expect(worker.ID).Should(Equal(""))
				Expect(worker.State).Should(Equal(""))
			})
		})
	})
	//Delete
	Describe("Delete", func() {
		Context("When delete of worker is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete cluster", func() {
				target := &ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newWorker(server.URL()).Delete("test", "abc-123-def-ghi", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When cluster delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete worker`),
					),
				)
			})

			It("should return error service key delete", func() {
				target := &ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newWorker(server.URL()).Delete("test", "abc-123-def-ghi", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Update
	Describe("Update", func() {
		Context("When update worker is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should return worker updated", func() {
				target := &ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WorkerParam{
					Action: "add", Count: 1,
				}
				err := newWorker(server.URL()).Update("test", "abc-123-def-ghi", params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When updating worker is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to add worker to cluster`),
					),
				)
			})

			It("should return error during updating worker", func() {
				target := &ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WorkerParam{
					Action: "add", Count: 1,
				}
				err := newWorker(server.URL()).Update("test", "abc-123-def-ghi", params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func newWorker(url string) Workers {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.CfService,
	}
	return newWorkerAPI(&client)
}
