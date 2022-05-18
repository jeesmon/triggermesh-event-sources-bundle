package main

import (
	"fmt"
	"log"
	"os"

	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"

	// AWS Event Sources

	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/awss3source"
)

func main() {
	ctx := signals.NewContext()

	ns, err := getWatchNamespace()
	if err != nil {
		log.Fatal("Unable to get WATCH_NAMESPACE")
	}

	ctx = injection.WithNamespaceScope(ctx, ns)

	sharedmain.MainWithContext(ctx, "triggermesh-controller",
		awss3source.NewController,
	)
}

// getWatchNamespace returns the Namespace the controller should be watching for changes
func getWatchNamespace() (string, error) {
	// WatchNamespaceEnvVar is the constant for env variable WATCH_NAMESPACE
	// which specifies the Namespace to watch.
	// An empty value means the operator is running with cluster scope.
	var watchNamespaceEnvVar = "WATCH_NAMESPACE"

	ns, found := os.LookupEnv(watchNamespaceEnvVar)
	if !found {
		return "", fmt.Errorf("%s must be set", watchNamespaceEnvVar)
	}
	return ns, nil
}
