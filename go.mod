module github.com/crossplane/provider-nop

go 1.13

require (
	github.com/crossplane/crossplane v0.11.0
	github.com/crossplane/crossplane-runtime v0.9.0
	github.com/crossplane/crossplane-tools v0.0.0-20200412230150-efd0edd4565b
	github.com/crossplane/provider-gcp v0.10.0 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/pkg/errors v0.8.1
	google.golang.org/api v0.21.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/controller-tools v0.2.4
)
