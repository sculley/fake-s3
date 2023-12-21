# Fake S3

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/sculley/fake-s3">
    <img src="https://raw.githubusercontent.com/sculley/fake-s3/main/images/logo.png" alt="Logo" width="350" height="350">
  </a>

<h3 align="center">Fake S3</h3>

  <p align="center">
    fake-s3 is a lightweight server that responds to the same API as Amazon S3.
    <br />
    <a href="https://github.com/sculley/fake-s3"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/sculley/fake-s3/issues">Report Bug/Issue</a>
    ·
    <a href="https://github.com/sculley/fake-s3/pulls">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->

<summary>Table of Contents</summary>
<ol>
  <li>
    <a href="#about-the-project">About The Project</a>
    <ul>
      <li><a href="#built-with">Built With</a></li>
    </ul>
  </li>
  <li>
    <a href="#getting-started">Getting Started</a>
    <ul>
      <li><a href="#installation">Installation</a></li>
    </ul>
  </li>
  <li><a href="#usage">Usage</a></li>
  <li><a href="#roadmap">Roadmap</a></li>
  <li><a href="#contributing">Contributing</a></li>
  <li><a href="#license">License</a></li>
  <li><a href="#contact">Contact</a></li>
  <li><a href="#acknowledgments">Acknowledgments</a></li>
</ol>


<!-- ABOUT THE PROJECT -->
## About The Project

fake-s3 is a lightweight server that responds to the same API as Amazon S3. It is extremely useful for testing S3 in a sandbox environment without actually making calls to Amazon, which is useful for testing code using S3 but without the cost and complications of using the real Amazon S3.

fake-s3 is built using the [Go-Fake-S3] library.

Please don't use Fake S3 as a production service. The intended use case for Fake S3 is currently to facilitate testing. It's not meant to be used for safe, persistent access to production data at the moment.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* [![Go][Go-Badge]][Go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

To run the project locally, you will need to make sure you have configured the `config.yml` correctly with the correct values for your environment. The `config.yml` file is located in the `conf` directory.

```yaml
general:
  port: 8080
  read_timeout: 15
  write_timeout: 15
```

### Running Locally

You can run the `Fake S3` as a docker container using the following commands;

* Build the container image

  ```sh
  docker build -t fake-s3:develop .
  ```

* Populate the `config.yaml` file with the various values.

* Run the container

  ```sh
  docker run -e CONFIG_PATH="./conf/" -p 8080:8080 -d --name fake-s3 fake-s3:develop
  ```

Or you can run the `Fake S3` as a standalone binary using the following commands;

* Build the binary

  ```sh
  go mod tidy
  go build -o bin/fake-s3 cmd/fake-s3/main.go
  ```

* Populate the `config.yml` file with the various values.

You can use the config.yml included in the repo as a template.

* Run the binary

  ```sh
  export CONFIG_PATH="./conf/"
  ./bin/fake-s3
  ```

### Usage

```go
// Setup a new config
cfg, _ := config.LoadDefaultConfig(
	context.TODO(),
    config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("KEY", "SECRET", "SESSION")),
    config.WithHTTPClient(&http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
    }),
    config.WithEndpointResolverWithOptions(
        aws.EndpointResolverWithOptionsFunc(func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
            return aws.Endpoint{URL: ts.URL}, nil
        }),
    ),
)

// Create an Amazon S3 v2 client, important to use o.UsePathStyle
// alternatively change local DNS settings, e.g., in /etc/hosts
// to support requests to http://<bucketname>.127.0.0.1:32947/...
client := s3.NewFromConfig(cfg, func(o *s3.Options) {
	o.UsePathStyle = true
})

// Create a bucket
client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
  Bucket: aws.String("fake-s3-bucket"),
})

// Upload an object
tempFile, err := os.CreateTemp(os.TempDir(), "test.txt")
if err != nil {
	fmt.Println(err)
}

result, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
	Bucket: aws.String("fake-s3-bucket"),
	Key:    aws.String(filepath.Base(tempFile.Name())),
	Body:   tempFile,
})
if err != nil {
	fmt.Println(err)
}
```

<!-- ROADMAP -->
## Issues

See the [open issues](https://github.com/sculley/fake-s3/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Clone or Fork the project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request to merge into develop

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the Apache License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Project Link: [https://github.com/sculley/fake-s3](https://github.com/sculley/fake-s3)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Go-Fake-S3]

<!-- MARKDOWN LINKS & IMAGES -->
[Go-Badge]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev
[Go-Fake-S3]: https://github.com/johannesboyne/gofakes3
