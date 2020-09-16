package main

import dsock "github.com/Cretezy/dSock-go"

func main() {
	dSockClient := dsock.NewClient("https://api.dsock.cloud/704f643b-1607-4cc9-bf4f-725221dc0a63", "7TyUQCuMWzuq19Aga7FF547IpA424XtF")

	claim, err := dSockClient.CreateClaim(dsock.CreateClaimOptions{
		User: "test",
	})

	if err != nil {
		panic(err)
	}

	println(claim.Id)

}
