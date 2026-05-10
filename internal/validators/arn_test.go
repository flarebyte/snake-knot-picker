// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package validators

import "testing"

func TestValidateARN(t *testing.T) {
	opts := ARNOptions{
		AllowPartition: []string{"aws"},
		AllowService:   []string{"s3", "sns"},
		AllowRegion:    []string{"us-east-2"},
		AllowAccountID: []string{"123456789012"},
		AllowResource:  []string{"example-sns-topic-name"},
	}
	ok := "arn:aws:sns:us-east-2:123456789012:example-sns-topic-name"
	if err := ValidateARN(ok, opts); err != nil {
		t.Fatalf("unexpected valid arn error: %v", err)
	}
	if err := ValidateARN("arn:aws:sqs:us-east-2:123456789012:example-sns-topic-name", opts); err == nil {
		t.Fatal("expected service allow-list failure")
	}
	if err := ValidateARN("arn:aws:sns:us-east-1:123456789012:example-sns-topic-name", opts); err == nil {
		t.Fatal("expected region allow-list failure")
	}
	if err := ValidateARN("arn:aws:sns:us-east-2:000000000000:example-sns-topic-name", opts); err == nil {
		t.Fatal("expected account allow-list failure")
	}
	if err := ValidateARN("arn:aws-cn:sns:us-east-2:123456789012:example-sns-topic-name", opts); err == nil {
		t.Fatal("expected partition allow-list failure")
	}
	if err := ValidateARN("arn:aws:sns:us-east-2:123456789012:other", opts); err == nil {
		t.Fatal("expected resource allow-list failure")
	}
	if err := ValidateARN("arn:aws:sns", opts); err == nil {
		t.Fatal("expected malformed arn failure")
	}
}

func TestValidateARNAdditionalBranches(t *testing.T) {
	// Empty allow-lists should allow valid ARN parts.
	open := ARNOptions{}
	if err := ValidateARN("arn:aws:s3:::bucket-name", open); err != nil {
		t.Fatalf("unexpected open allow-list rejection: %v", err)
	}
	// Malformed: missing service.
	if err := ValidateARN("arn:aws::us-east-1:123456789012:res", open); err == nil {
		t.Fatal("expected malformed missing-service rejection")
	}
	// Malformed: missing resource.
	if err := ValidateARN("arn:aws:sns:us-east-1:123456789012:", open); err == nil {
		t.Fatal("expected malformed missing-resource rejection")
	}
}
