.PHONY: run, sol

run:
	go run main.go

sol:
	solc --optimize --abi ./contracts/PunchFaceContract.sol -o build
	solc --optimize --bin ./contracts/PunchFaceContract.sol -o build
	abigen --abi=./build/PunchFaceContract.abi --bin=./build/PunchFaceContract.bin --pkg=punchface --out=./punchface/PunchFaceContract.go
