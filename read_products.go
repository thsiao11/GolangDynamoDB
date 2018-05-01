/*

This program performs a simple task of connecting to the DynamoDB test.product 
table at region: us-west-1 on Amazon AWS Web Service, query and display 
relevant product information sorted by price in descending order.

*/

package main

import (
    "fmt"
    "os"
    "sort"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// Create structs to hold Product and related fields
type Product struct {
    Prod_id int`json:"id"`
    Prod_description string`json:"description"`
    Unit string`json:"unit"`
    Price float64`json:"price"`
}

func main() {

    // create an array of products to be used for sorting by price
    var result_arr []Product

    // Initialize a session in us-west-1 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-1"),
        MaxRetries: aws.Int(3)},
    )

    if err != nil {
        fmt.Println("Error while creating session:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    // Create a new DynamoDB client
    svc := dynamodb.New(sess)

    // Project to Product ID, description, and price
    proj := expression.NamesList(expression.Name("id"), 
        expression.Name("description"), expression.Name("unit"),
        expression.Name("price"))

    expr, err := expression.NewBuilder().WithProjection(proj).Build()

    if err != nil {
        fmt.Println("Errors while building expression:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    // Build the query input parameters
    params := &dynamodb.ScanInput{
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
        ProjectionExpression:      expr.Projection(),
        // Product Table named: test.product
        TableName:                 aws.String("test.product"),
    }

    // Call DynamoDB Query API
    result, err := svc.Scan(params)

    if err != nil {
        fmt.Println("Query API call failed:")
        fmt.Println((err.Error()))
        os.Exit(1)
    }

    for _, v := range result.Items {
        item := Product{}

        err = dynamodbattribute.UnmarshalMap(v, &item)

        if err != nil {
            fmt.Println("Error when unmarshalling data:")
            fmt.Println(err.Error())
            os.Exit(1)
        }

        result_arr = append(result_arr, item)
    }

    // Sort result_arr by Price in descending order
    sort.Slice(result_arr, func(i, j int) bool { return result_arr[i].Price > result_arr[j].Price })

    // Print out each product in sorted order
    for _, v := range result_arr {
        fmt.Println("Product ID: ", v.Prod_id)
        fmt.Println("Product Description: ", v.Prod_description)
        fmt.Println("Unit: ", v.Unit)
        fmt.Println("Price: ", v.Price)
        fmt.Println()        
    }
}
