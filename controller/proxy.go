package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"goProxy/config"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

//var proxy = &httputil.ProxyRequest{
//	Director: func(req *http.Request) {
//		u, _ := url.Parse("https://oaistatic-cdn.closeai.biz")
//		req.URL = u
//		req.Host = u.Host
//
//	},
//}

var proxy = &httputil.ReverseProxy{
	Director: func(req *http.Request) {
		proxyUrl, _ := url.Parse(config.ProxyUrl)
		req.URL.Scheme = proxyUrl.Scheme
		req.URL.Host = proxyUrl.Host
		req.URL.Path = proxyUrl.Path + req.URL.Path
		req.Host = proxyUrl.Host
		req.Header.Del("Accept-Encoding")
	},
	ModifyResponse: func(res *http.Response) error {
		if res.StatusCode >= 400 {
			// g.Log().Warning(ctx, "ProxyNext:", response.StatusCode, response.Request.Method, response.Request.URL.String())
			config.Log.Warn("ProxyNext:", res.StatusCode, res.Request.Method, res.Request.URL.String())
		} else {
			// g.Log().Info(ctx, "ProxyNext:", response.StatusCode, response.Request.Method, response.Request.URL.String())
			config.Log.Info("ProxyNext:", res.StatusCode, res.Request.Method, res.Request.URL.String())
		}

		if res.StatusCode == 200 && strings.Contains(res.Header.Get("Content-Type"), "text") {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				// g.Log().Error(ctx, err)
				config.Log.Error(err)
				return err
			}

			proxyUrl, _ := url.Parse(config.ProxyUrl)
			currentUrl, _ := url.Parse(config.CurrentUrl)
			body = bytes.Replace(body, []byte(proxyUrl.Host), []byte(currentUrl.Host), -1)
			res.Body = io.NopCloser(strings.NewReader(string(body)))
		}
		return nil
	},
}

func ProxyAll(c *gin.Context) {
	//var proxy = &httputil.ReverseProxy{
	//	Director: func(req *http.Request) {
	//
	//		proxyUrl, _ := url.Parse(config.ProxyUrl)
	//		req.URL.Scheme = proxyUrl.Scheme
	//		req.URL.Host = proxyUrl.Host
	//		req.URL.Path = proxyUrl.Path + req.URL.Path
	//		req.Host = proxyUrl.Host
	//		req.Header.Del("Accept-Encoding")
	//	},
	//	ModifyResponse: func(res *http.Response) error {
	//		if res.StatusCode >= 400 {
	//			// g.Log().Warning(ctx, "ProxyNext:", response.StatusCode, response.Request.Method, response.Request.URL.String())
	//			config.Log.Warn("ProxyNext:", res.StatusCode, res.Request.Method, res.Request.URL.String())
	//		} else {
	//			// g.Log().Info(ctx, "ProxyNext:", response.StatusCode, response.Request.Method, response.Request.URL.String())
	//			config.Log.Info("ProxyNext:", res.StatusCode, res.Request.Method, res.Request.URL.String())
	//		}
	//
	//		if res.StatusCode == 200 && strings.Contains(res.Header.Get("Content-Type"), "text") {
	//			body, err := io.ReadAll(res.Body)
	//			if err != nil {
	//				// g.Log().Error(ctx, err)
	//				config.Log.Error(err)
	//				return err
	//			}
	//			body = bytes.Replace(body, []byte("oaistatic-cdn.closeai.biz"), []byte("oai-static.hicafes.com"), -1)
	//			res.Body = io.NopCloser(strings.NewReader(string(body)))
	//		}
	//		return nil
	//	},
	//}
	proxy.ServeHTTP(c.Writer, c.Request)
}
