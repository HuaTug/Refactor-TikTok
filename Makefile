idl:
	./kitex_gen.sh
env:
	docker-compose up -d
build:
	cd cmd/api && go build -o api
	cd cmd/user && go build -o user
	cd cmd/video && go build -o video
	cd cmd/relation && go build -o relation	
	cd cmd/interaction && go build -o interaction
	cd cmd/message && go build -o message 
api:
	cd cmd/api && ./api
user:	
	cd cmd/user &&./user
video:
	cd cmd/video &&./video
interaction:
	cd cmd/interaction && ./interaction
relation:
	cd cmd/relation && ./relation		
message:
	cd cmd/message && ./message

go:env build

clean:
	rm -f message