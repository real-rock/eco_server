conda activate economicus

protoc -I./ --go_out=../../main/grpc/proto --go_opt=paths=source_relative \
--go-grpc_out=../../main/grpc/proto --go-grpc_opt=paths=source_relative quant.proto

python -m grpc_tools.protoc -I. \
 --python_out=../../quant/proto \
 --grpc_python_out=../../quant/proto quant.proto

conda deactivate
