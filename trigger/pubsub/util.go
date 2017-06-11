package pubsub

import (
	"strings"
)

// "gcr.io/v2-namespace/hello-world:1.1"
func extractContainerRegistryURI(imageName string) string {
	parts := strings.Split(imageName, "/")
	return parts[0]
}

func containerRegistryURI(projectID, registry string) string {
	return registry + "%2F" + projectID
}

func containerRegistrySubName(projectID, topic string) string {
	return "keel-" + projectID + "-" + topic
}

// isGoogleContainerRegistry - we only care about gcr.io images,
// with other registries - we won't be able to receive events.
// Theoretically if someone publishes messages for updated images to
// google pubsub - we could turn this off
func isGoogleContainerRegistry(registry string) bool {
	return strings.Contains(registry, "gcr.io")
}
