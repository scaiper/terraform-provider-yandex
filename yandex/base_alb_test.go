package yandex

import (
	"fmt"
	"os"

	"github.com/fatih/structs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1"
)

const albDefaultSni = "tf-test-tls"
const albDefaultValidationContext = "tf-test-validation-context"
const albDefaultBackendWeight = "1"
const albDefaultPanicThreshold = "50"
const albDefaultLocalityPercent = "35"
const albDefaultTimeout = "3s"
const albDefaultInterval = "5s"
const albDefaultStrictLocality = "true"
const albDefaultServiceName = "true"
const albDefaultHTTP2 = "true"
const albDefaultHost = "tf-test-host"
const albDefaultPath = "tf-test-path"
const albDefaultPort = "3"
const albDefaultSend = "tf-test-send"
const albDefaultReceive = "tf-test-receive"
const albDefaultDescription = "alb-bg-description"
const albDefaultDirectResponseBody = "Not Found"
const albDefaultDirectResponseStatus = "404"
const albDefaultModificationAppend = "header"
const albDefaultStatusResponse = "not_found"
const albDefaultRedirectReplacePrefix = "/other"
const albDefaultAutoHostRewrite = "true"
const albDefaultAllowHTTP10 = "true"
const albDefaultMaxConcurrentStreams = "2"
const albDefaultHTTPToHTTPS = "true"

type resourceALBLoadBalancerInfo struct {
	IsHTTPListener        bool
	IsTLSListener         bool
	IsRedirects           bool
	IsHTTPListenerHandler bool
	IsTLSListenerHandler  bool
	IsDataSource          bool
	IsHTTP2Options        bool
	IsAllowHTTP10         bool

	BaseTemplate string

	BalancerName         string
	RouterName           string
	ListenerName         string
	BalancerDescription  string
	AllowHTTP10          string
	MaxConcurrentStreams string
	EndpointPort         string
	HTTPToHTTPS          string
	CertificateID        string
}

func albLoadBalancerInfo() resourceALBLoadBalancerInfo {
	res := resourceALBLoadBalancerInfo{
		IsHTTPListener:        false,
		IsTLSListener:         false,
		IsDataSource:          false,
		IsRedirects:           false,
		IsHTTPListenerHandler: false,
		IsTLSListenerHandler:  false,
		IsHTTP2Options:        false,
		IsAllowHTTP10:         false,
		BaseTemplate:          testAccALBBaseTemplate(acctest.RandomWithPrefix("tf-instance")),
		BalancerName:          acctest.RandomWithPrefix("tf-load-balancer"),
		RouterName:            acctest.RandomWithPrefix("tf-router"),
		ListenerName:          acctest.RandomWithPrefix("tf-listener"),
		BalancerDescription:   acctest.RandomWithPrefix("tf-load-balancer-description"),
		AllowHTTP10:           albDefaultAllowHTTP10,
		MaxConcurrentStreams:  albDefaultMaxConcurrentStreams,
		EndpointPort:          albDefaultPort,
		HTTPToHTTPS:           albDefaultHTTPToHTTPS,
		CertificateID:         os.Getenv("ALB_TEST_CERTIFICATE_ID"),
	}

	return res
}

type resourceALBVirtualHostInfo struct {
	IsModifyRequestHeaders     bool
	IsModifyResponseHeaders    bool
	IsHTTPRoute                bool
	IsGRPCRoute                bool
	IsHTTPRouteAction          bool
	IsRedirectAction           bool
	IsDirectResponseAction     bool
	IsGRPCRouteAction          bool
	IsGRPCStatusResponseAction bool
	IsDataSource               bool

	BaseTemplate string

	VHName     string
	TGName     string
	RouterName string
	BGName     string

	RouterDescription              string
	ModificationAppend             string
	RouteName                      string
	DirectResponseStatus           string
	DirectResponseBody             string
	RedirectReplacePrefix          string
	HTTPRouteActionTimeout         string
	GRPCRouteActionTimeout         string
	GRPCStatusResponseActionStatus string
	GRPCRouteActionAutoHostRewrite string
}

func albVirtualHostInfo() resourceALBVirtualHostInfo {
	res := resourceALBVirtualHostInfo{
		IsModifyRequestHeaders:         false,
		IsModifyResponseHeaders:        false,
		IsHTTPRoute:                    false,
		IsGRPCRoute:                    false,
		IsHTTPRouteAction:              false,
		IsRedirectAction:               false,
		IsDirectResponseAction:         false,
		IsGRPCRouteAction:              false,
		IsGRPCStatusResponseAction:     false,
		IsDataSource:                   false,
		BaseTemplate:                   testAccALBBaseTemplate(acctest.RandomWithPrefix("tf-instance")),
		VHName:                         acctest.RandomWithPrefix("tf-virtual-host"),
		TGName:                         acctest.RandomWithPrefix("tf-tg"),
		RouterName:                     acctest.RandomWithPrefix("tf-router"),
		BGName:                         acctest.RandomWithPrefix("tf-bg"),
		RouterDescription:              albDefaultDescription,
		RouteName:                      acctest.RandomWithPrefix("tf-route"),
		ModificationAppend:             albDefaultModificationAppend,
		DirectResponseStatus:           albDefaultDirectResponseStatus,
		DirectResponseBody:             albDefaultDirectResponseBody,
		RedirectReplacePrefix:          albDefaultRedirectReplacePrefix,
		HTTPRouteActionTimeout:         albDefaultTimeout,
		GRPCRouteActionTimeout:         albDefaultTimeout,
		GRPCStatusResponseActionStatus: albDefaultStatusResponse,
		GRPCRouteActionAutoHostRewrite: albDefaultAutoHostRewrite,
	}

	return res
}

type resourceALBBackendGroupInfo struct {
	IsHTTPBackend bool
	IsGRPCBackend bool
	IsHTTPCheck   bool
	IsGRPCCheck   bool
	IsStreamCheck bool
	IsDataSource  bool
	IsEmptyTLS    bool

	BaseTemplate string

	TGName string
	BGName string

	BGDescription        string
	TlsSni               string
	TlsValidationContext string
	BackendWeight        string
	PanicThreshold       string
	LocalityPercent      string
	StrictLocality       string
	Timeout              string
	Interval             string
	ServiceName          string
	HTTP2                string
	Host                 string
	Path                 string
	Port                 string
	Receive              string
	Send                 string
}

func albBackendGroupInfo() resourceALBBackendGroupInfo {
	res := resourceALBBackendGroupInfo{
		IsHTTPBackend:        false,
		IsGRPCBackend:        false,
		IsHTTPCheck:          false,
		IsGRPCCheck:          false,
		IsStreamCheck:        false,
		IsDataSource:         false,
		IsEmptyTLS:           false,
		BaseTemplate:         testAccALBBaseTemplate(acctest.RandomWithPrefix("tf-instance")),
		TGName:               acctest.RandomWithPrefix("tf-tg"),
		BGName:               acctest.RandomWithPrefix("tf-bg"),
		BGDescription:        albDefaultDescription,
		TlsSni:               albDefaultSni,
		TlsValidationContext: albDefaultValidationContext,
		BackendWeight:        albDefaultBackendWeight,
		PanicThreshold:       albDefaultPanicThreshold,
		LocalityPercent:      albDefaultLocalityPercent,
		StrictLocality:       albDefaultStrictLocality,
		Timeout:              albDefaultTimeout,
		Interval:             albDefaultInterval,
		ServiceName:          albDefaultServiceName,
		HTTP2:                albDefaultHTTP2,
		Host:                 albDefaultHost,
		Path:                 albDefaultPath,
		Port:                 albDefaultPort,
		Receive:              albDefaultReceive,
		Send:                 albDefaultSend,
	}

	return res
}

const albVirtualHostConfigTemplate = `
{{ if .IsDataSource }}
data "yandex_alb_virtual_host" "test-virtual-host-ds" {
  virtual_host_id = yandex_alb_virtual_host.test-vh.id
}		
{{ end }}
resource "yandex_alb_http_router" "test-router" {
  name        = "{{.RouterName}}"
  description = "{{.RouterDescription}}"
}
resource "yandex_alb_backend_group" "test-bg" {
  name        = "{{.BGName}}"
  {{if .IsHTTPRoute}}
  http_backend {
    name             = "test-http-backend"
    weight           = 1
    target_group_ids = ["${yandex_alb_target_group.test-target-group.id}"]
  }
  {{end}}
  {{if .IsGRPCRoute}}
  grpc_backend {
    name             = "test-grpc-backend"
    weight           = 1
    target_group_ids = ["${yandex_alb_target_group.test-target-group.id}"]
  }
  {{end}}
}
resource "yandex_alb_virtual_host" "test-vh" {
  name        = "{{.VHName}}"
  http_router_id = yandex_alb_http_router.test-router.id

  authority = ["*.foo.com", "*-bar.foo.com"]

  modify_request_headers {
    name   = "test-header"
    append = "{{.ModificationAppend}}"
  } 
  {{ if or .IsHTTPRoute .IsGRPCRoute}}
  route {
    name = "{{.RouteName}}"
    {{if .IsHTTPRoute}}
    http_route {
      http_match {
        path {
          prefix = "/http/match/"
        }
        http_method = ["GET", "PUT"]
      }
      {{if .IsHTTPRouteAction}}
      http_route_action {
        backend_group_id = yandex_alb_backend_group.test-bg.id
        timeout = "{{.HTTPRouteActionTimeout}}"
      }
      {{end}}
      {{if .IsDirectResponseAction}}
      direct_response_action {
        status = {{.DirectResponseStatus}}
        body = "{{.DirectResponseBody}}"
      }  
      {{end}}
      {{if .IsRedirectAction}}
      redirect_action {
        replace_prefix = "{{.RedirectReplacePrefix}}"
      }
      {{end}}
    }
    {{end}}
    {{if .IsGRPCRoute}}
    grpc_route {
      grpc_match {
        fqmn {
          exact = "some.service"
        }
      }
      {{if .IsGRPCRouteAction}}
      grpc_route_action {
        backend_group_id = yandex_alb_backend_group.test-bg.id
        max_timeout = "{{.GRPCRouteActionTimeout}}"
        auto_host_rewrite = {{.GRPCRouteActionAutoHostRewrite}}
      }
      {{end}}
      {{if .IsGRPCStatusResponseAction}}
      grpc_status_response_action {
        status = "{{.GRPCStatusResponseActionStatus}}"
      }  
      {{end}}
    }
    {{end}}
  }
  {{end}}
}
{{ if or .IsHTTPRoute .IsGRPCRoute }}
resource "yandex_alb_target_group" "test-target-group" {
  name		= "{{.TGName}}"

  target {
	subnet_id	= "${yandex_vpc_subnet.test-subnet.id}"
	ip_address	= "${yandex_compute_instance.test-instance-1.network_interface.0.ip_address}"
  }

  target {
	subnet_id	= "${yandex_vpc_subnet.test-subnet.id}"
	ip_address	= "${yandex_compute_instance.test-instance-2.network_interface.0.ip_address}"
  }
}
{{ end }}
{{.BaseTemplate}}
`

const albLoadBalancerConfigTemplate = `
{{ if .IsDataSource }}
data "yandex_alb_load_balancer" "test-alb-ds" {
  name = yandex_alb_load_balancer.test-balancer.name
}
{{ end }}
resource "yandex_alb_http_router" "test-router" {
  name        = "{{.RouterName}}"
}

resource "yandex_alb_load_balancer" "test-balancer" {
  name        = "{{.BalancerName}}"
  description = "{{.BalancerDescription}}"

  security_group_ids = [yandex_vpc_security_group.test-security-group.id]
  network_id  = yandex_vpc_network.test-network.id
  labels = {
    tf-label    = "tf-label-value"
    empty-label = ""
  }
  allocation_policy {
    location {
      zone_id   = "ru-central1-a"
      subnet_id = yandex_vpc_subnet.test-subnet.id 
    }
  }
  {{ if or .IsHTTPListener .IsTLSListener }}
  listener {
    name = "{{.ListenerName}}"
    endpoint {
      address {
        internal_ipv4_address {
          subnet_id = yandex_vpc_subnet.test-subnet.id
        }
      }
      ports = [ {{.EndpointPort}} ]
    }    
    {{if .IsHTTPListener}}
    http {
      {{if .IsHTTPListenerHandler}}
      handler {
        http_router_id = yandex_alb_http_router.test-router.id
        {{if .IsAllowHTTP10}}
        allow_http10 = {{.AllowHTTP10}}
        {{end}}
        {{if .IsHTTP2Options}}
        http2_options {
          max_concurrent_streams = {{.MaxConcurrentStreams}}
        }
        {{end}}
      }
      {{end}}
      {{if .IsRedirects}}
      redirects {
        http_to_https = {{.HTTPToHTTPS}}
      }
      {{end}}
    }
    {{end}}
    {{if .IsTLSListener}}
    tls {
      default_handler {
        http_handler {
          http_router_id = yandex_alb_http_router.test-router.id
          {{if .IsAllowHTTP10}}
            allow_http10 = {{.AllowHTTP10}}
          {{end}}
          {{if .IsHTTP2Options}}
          http2_options {
            max_concurrent_streams = {{.MaxConcurrentStreams}}
          }
          {{end}}
        }
        certificate_ids = ["{{.CertificateID}}"]
      }
      sni_handler {
        name = "host"
        server_names = ["host.url.com"]
        handler {
          http_handler {
            http_router_id = yandex_alb_http_router.test-router.id
          }
          certificate_ids = ["{{.CertificateID}}"]
        }
      }
    }
    {{end}}
  }    
  {{end}}
}

resource "yandex_vpc_security_group" "test-security-group" {
  network_id = yandex_vpc_network.test-network.id

  ingress {
    protocol       = "TCP"
    description    = "healthchecks"
    port           = 30080
    v4_cidr_blocks = ["198.18.235.0/24", "198.18.248.0/24"]
  }
}

{{.BaseTemplate}}
`

const albBackendGroupConfigTemplate = `
{{ if .IsDataSource }}
data "yandex_alb_backend_group" "test-bg-ds" {
  name = yandex_alb_backend_group.test-bg.name
}		
{{ end }}
resource "yandex_alb_backend_group" "test-bg" {
  name        = "{{.BGName}}"
  description = "{{.BGDescription}}"

  labels = {
    tf-label    = "tf-label-value"
    empty-label = ""
  }
  {{ if .IsHTTPBackend }}
  http_backend {
    name             = "test-http-backend"
    weight           = {{.BackendWeight}}
    port             = {{.Port}}
    target_group_ids = ["${yandex_alb_target_group.test-target-group.id}"]
    tls {
      {{ if not .IsEmptyTLS }}
      sni = "{{.TlsSni}}"
      validation_context {
        trusted_ca_bytes = "{{.TlsValidationContext}}"
      }
      {{end}}
    }
    load_balancing_config {
      panic_threshold                = {{.PanicThreshold}}
      locality_aware_routing_percent = {{.LocalityPercent}}
      strict_locality                = {{.StrictLocality}}
    }
    {{ if .IsGRPCCheck }}
    healthcheck {
      timeout  = "{{.Timeout}}"
      interval = "{{.Interval}}"
      grpc_healthcheck {
        service_name = "{{.ServiceName}}"
      }
    }
    {{end}}
    {{ if .IsStreamCheck }}
    healthcheck {
      timeout  = "{{.Timeout}}"
      interval = "{{.Interval}}"
      stream_healthcheck {
        receive = "{{.Receive}}"
        send    = "{{.Send}}"
      }
    }
    {{end}}
    {{ if .IsHTTPCheck }}
    healthcheck {
      timeout = "{{.Timeout}}"
      interval = "{{.Interval}}"
      http_healthcheck {
        host  = "{{.Host}}"
        path  = "{{.Path}}"
        http2 = "{{.HTTP2}}"
      }
    }
    {{end}}
    http2 = "{{.HTTP2}}"
  }
  {{end}}
  {{ if .IsGRPCBackend }}
  grpc_backend {
    name             = "test-grpc-backend"
    weight           = {{.BackendWeight}}
    port             = {{.Port}}
    target_group_ids = ["${yandex_alb_target_group.test-target-group.id}"]
    tls {
      sni = "{{.TlsSni}}"
      validation_context {
        trusted_ca_bytes = "{{.TlsValidationContext}}"
      }
    }
    load_balancing_config {
      panic_threshold                = {{.PanicThreshold}}
      locality_aware_routing_percent = {{.LocalityPercent}}
      strict_locality                = {{.StrictLocality}}
    }
    {{ if .IsGRPCCheck }}
    healthcheck {
      timeout  = "{{.Timeout}}"
      interval = "{{.Interval}}"
      grpc_healthcheck {
        service_name = "{{.ServiceName}}"
      }
    }
    {{end}}
    {{ if .IsStreamCheck }}
    healthcheck {
      timeout  = "{{.Timeout}}"
      interval = "{{.Interval}}"
      stream_healthcheck {
        receive = "{{.Receive}}"
        send    = "{{.Send}}"
      }
    }
    {{end}}
    {{ if .IsHTTPCheck }}
    healthcheck {
      timeout  = "{{.Timeout}}"
      interval = "{{.Interval}}"
      http_healthcheck {
        host  = "{{.Host}}"
        path  = "{{.Path}}"
        http2 = "{{.HTTP2}}"
      }
    }
    {{end}}
  }
  {{end}}
}
{{ if or .IsHTTPBackend .IsGRPCBackend }}
resource "yandex_alb_target_group" "test-target-group" {
  name		= "{{.TGName}}"

  target {
	subnet_id	= "${yandex_vpc_subnet.test-subnet.id}"
	ip_address	= "${yandex_compute_instance.test-instance-1.network_interface.0.ip_address}"
  }

  target {
	subnet_id	= "${yandex_vpc_subnet.test-subnet.id}"
	ip_address	= "${yandex_compute_instance.test-instance-2.network_interface.0.ip_address}"
  }
}
{{ end }}

{{.BaseTemplate}}
`

func testALBBackendGroupConfig_basic(in resourceALBBackendGroupInfo) string {
	m := structs.Map(in)
	config := templateConfig(albBackendGroupConfigTemplate, m)
	return config
}

func testALBLoadBalancerConfig_basic(in resourceALBLoadBalancerInfo) string {
	m := structs.Map(in)
	config := templateConfig(albLoadBalancerConfigTemplate, m)
	return config
}

func testALBVirtualHostConfig_basic(in resourceALBVirtualHostInfo) string {
	m := structs.Map(in)
	config := templateConfig(albVirtualHostConfigTemplate, m)
	return config
}

func testAccCheckALBBackendGroupValues(bg *apploadbalancer.BackendGroup, expectedHTTPBackends, expectedGRPCBackends bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if (bg.GetHttp() != nil) != expectedHTTPBackends {
			return fmt.Errorf("invalid presence or absence of http backend Application Backend Group %s", bg.Name)
		}

		if (bg.GetGrpc() != nil) != expectedGRPCBackends {
			return fmt.Errorf("invalid presence or absence of grpc backend Application Backend Group %s", bg.Name)
		}

		return nil
	}
}

func testAccCheckALBVirtualHostValues(vh *apploadbalancer.VirtualHost, expectedHttpRoute, expectedGrpcRoute bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, route := range vh.GetRoutes() {
			if (route.GetHttp() != nil) != expectedHttpRoute {
				return fmt.Errorf("invalid presence or absence of http backend Application Backend Group %s", vh.Name)
			}

			if (route.GetGrpc() != nil) != expectedGrpcRoute {
				return fmt.Errorf("invalid presence or absence of grpc backend Application Backend Group %s", vh.Name)
			}
		}

		return nil
	}
}

func testAccALBGeneralTGTemplate(tgName, tgDesc, baseTemplate string, targetsCount int, isDataSource bool) string {
	targets := make([]string, targetsCount)
	for i := 1; i <= targetsCount; i++ {
		targets[i-1] = fmt.Sprintf("test-instance-%d", i)
	}
	return templateConfig(`
{{ if .IsDataSource }}
data "yandex_alb_target_group" "test-tg-ds" {
  name = yandex_alb_target_group.test-tg.name
}		
{{ end }}
resource "yandex_alb_target_group" "test-tg" {
  name        = "{{.TGName}}"
  description = "{{.TGDescription}}"

  labels = {
    tf-label    = "tf-label-value"
    empty-label = ""
  }
  {{range .Targets}}
  target {
    subnet_id  = yandex_vpc_subnet.test-subnet.id
    ip_address = yandex_compute_instance.{{.}}.network_interface.0.ip_address
  }		
  {{end}}
}

{{.BaseTemplate}}
		`,
		map[string]interface{}{
			"TGName":        tgName,
			"TGDescription": tgDesc,
			"BaseTemplate":  baseTemplate,
			"Targets":       targets,
			"IsDataSource":  isDataSource,
		},
	)
}

func testAccCheckALBTargetGroupValues(tg *apploadbalancer.TargetGroup, expectedInstanceNames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		subnetIPMap, err := getSubnetIPMap(expectedInstanceNames)
		if err != nil {
			return err
		}

		if len(tg.GetTargets()) != len(expectedInstanceNames) {
			return fmt.Errorf("invalid count of targets in Application Target Group %s", tg.Name)
		}

		for _, t := range tg.GetTargets() {
			if addresses, ok := subnetIPMap[t.GetSubnetId()]; ok {
				addressExists := false
				for _, a := range addresses {
					if a == t.GetIpAddress() {
						addressExists = true
						break
					}
				}
				if !addressExists {
					return fmt.Errorf("invalid Target's Address %s in Application Target Group %s", t.GetIpAddress(), tg.Name)
				}
			} else {
				return fmt.Errorf("invalid Target's SubnetID %s in Application Target Group %s", t.GetSubnetId(), tg.Name)
			}
		}

		return nil
	}
}

func testAccALBBaseTemplate(instanceName string) string {
	return fmt.Sprintf(`
data "yandex_compute_image" "test-image" {
  family = "ubuntu-1804-lts"
}

resource "yandex_compute_instance" "test-instance-1" {
  name        = "%[1]s-1"
  platform_id = "standard-v2"
  zone        = "ru-central1-a"

  resources {
    cores         = 2
    core_fraction = 20
    memory        = 2
  }

  boot_disk {
    initialize_params {
      size     = 4
      image_id = data.yandex_compute_image.test-image.id
    }
  }

  network_interface {
    subnet_id = yandex_vpc_subnet.test-subnet.id
  }

  scheduling_policy {
    preemptible = true
  }
}

resource "yandex_compute_instance" "test-instance-2" {
  name        = "%[1]s-2"
  platform_id = "standard-v2"
  zone        = "ru-central1-a"

  resources {
    cores         = 2
    core_fraction = 20
    memory        = 2
  }

  boot_disk {
    initialize_params {
      size     = 4
      image_id = data.yandex_compute_image.test-image.id
    }
  }

  network_interface {
    subnet_id = yandex_vpc_subnet.test-subnet.id
  }

  scheduling_policy {
    preemptible = true
  }
}

resource "yandex_vpc_network" "test-network" {}

resource "yandex_vpc_subnet" "test-subnet" {
  zone           = "ru-central1-a"
  network_id     = yandex_vpc_network.test-network.id
  v4_cidr_blocks = ["192.168.0.0/24"]
}
`, instanceName,
	)
}
