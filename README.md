This program performs a simple task of connecting to the DynamoDB test.product table at region: us-west-1 on Amazon AWS Web Service, query and display the relevant product information sorted by price in descending order.


Included:
Documentation: README
Source files: read_products.go

Requirements:
- Compile program with Go version 1.8+
- Program uses shared access credentials stored in the ~/.aws/credentials file

Assumptions:
- This is a test batch client to connect to the AWS Web Service DynamoDB database
- For product table with many attributes, build and use Global Secondary Index(GSI) to selectively extracts specific attributes to improve query performance and response time

If this program is used for as a server or real-time application, add the following features:
- Save product result data to memcache after initial database extraction
- Enable Autoscaling
- Implement connection pooling
- Use Global Secondary Index(GSI) for product table with large number of attributes
