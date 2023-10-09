## Go mod 将一些常用的包 replace

        replace(
            golang.org/x/text => github.com/golang/text latest
            golang.org/x/net => github.com/golang/net latest
            golang.org/x/crypto => github.com/golang/crypto latest
            golang.org/x/tools => github.com/golang/tools latest
            golang.org/x/sync => github.com/golang/sync latest
            golang.org/x/sys => github.com/golang/sys latest
            cloud.google.com/go => github.com/googleapis/google-cloud-go latest
            google.golang.org/genproto => github.com/google/go-genproto latest
            golang.org/x/exp => github.com/golang/exp latest
            golang.org/x/time => github.com/golang/time latest
            golang.org/x/oauth2 => github.com/golang/oauth2 latest
            golang.org/x/lint => github.com/golang/lint latest
            google.golang.org/grpc => github.com/grpc/grpc-go latest
            google.golang.org/api => github.com/googleapis/google-api-go-client latest
            google.golang.org/appengine => github.com/golang/appengine latest
            golang.org/x/mobile => github.com/golang/mobile latest
            golang.org/x/image => github.com/golang/image latest
        )