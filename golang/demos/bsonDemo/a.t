[32m.[0m[32m.[0m[33mx[0m[31m[0m[33m
Failures:

  * /Users/Jialin/golang/src/demo/bsonDemo/mongoDemo_test.go 
  Line 20:
  Expected     '<nil>'
  to NOT equal '<nil>'
  (but it did)!

[0m[33m
3 assertions thus far[0m

--- FAIL: TestStructTagJsonUnmarshalDemo (0.00s)
	mongoDemo_test.go:15: json ret: {"name":"jialin","phone":"p","Id":101}
	mongoDemo_test.go:18: uj: &{Name:jialin Phone:p Id:101}
[32m.[0m[32m.[0m[33mx[0m[31m[0m[33m
Failures:

  * /Users/Jialin/golang/src/demo/bsonDemo/mongoDemo_test.go 
  Line 35:
  Expected     '<nil>'
  to NOT equal '<nil>'
  (but it did)!

[0m[33m
6 assertions thus far[0m

--- FAIL: TestStructTagBsonUnmarshalDemo (0.00s)
	mongoDemo_test.go:33: &{  0}
[33mx[0m[31m[0m[33m
Failures:

  * /Users/Jialin/golang/src/demo/bsonDemo/mongoDemo_test.go 
  Line 46:
  Expected     '<nil>'
  to NOT equal '<nil>'
  (but it did)!

[0m[33m
7 assertions thus far[0m

--- FAIL: TestStructTagBsonDemo (0.00s)
	mongoDemo_test.go:44: [43 0 0 0 2 78 97 109 101 0 7 0 0 0 106 105 97 108 105 110 0 2 80 104 111 110 101 0 2 0 0 0 112 0 16 73 100 0 101 0 0 0 0]
	mongoDemo_test.go:45: ret:+   Name    jialin Phone    p Id e    
[33mx[0m[31m[0m[33m
Failures:

  * /Users/Jialin/golang/src/demo/bsonDemo/mongoDemo_test.go 
  Line 58:
  Expected     '<nil>'
  to NOT equal '<nil>'
  (but it did)!

[0m[33m
8 assertions thus far[0m

--- FAIL: TestStructTagJsonDemo (0.00s)
	mongoDemo_test.go:56: [123 34 110 97 109 101 34 58 34 106 105 97 108 105 110 34 44 34 112 104 111 110 101 34 58 34 112 34 44 34 73 100 34 58 49 125]
	mongoDemo_test.go:57: ret:{"name":"jialin","phone":"p","Id":1}
FAIL
exit status 1
FAIL	demo/bsonDemo	0.006s
