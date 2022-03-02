import datetime
import json
from concurrent import futures

import grpc
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToJson

from proto import quant_pb2, quant_pb2_grpc
from calculator.calculate import *
from datetime import datetime, timezone


class Quant(quant_pb2_grpc.QuantServicer):
    def Request(self, request, context):
        print('======= Got request =======')
        json_obj = MessageToJson(request, including_default_value_fields=True)
        req = json.loads(json_obj)

        date_format = "%Y-%m-%dT%H:%M:%SZ"
        if req["start_date"] == '0001-01-01T00:00:00Z':
            req["start_date"] = None
        else:
            req["start_date"] = datetime.strptime(req["start_date"], date_format).replace(tzinfo=timezone.utc)

        if req["end_date"] == '0001-01-01T00:00:00Z':
            req["end_date"] = None
        else:
            req["end_date"] = datetime.strptime(req["end_date"], date_format).replace(tzinfo=timezone.utc)

        req["net_revenue"] = {'min': int(req["net_revenue"]["min"]), 'max': int(req["net_revenue"]["max"])}
        req["net_profit"] = {'min': int(req["net_profit"]["min"]), 'max': int(req["net_profit"]["max"])}
        req["market_cap"] = {'min': int(req["market_cap"]["min"]), 'max': int(req["market_cap"]["max"])}

        print("request: ", req)
        try:
            calc = QuantCalc()
            result = calc.execute(
                start_date=req["start_date"],
                end_date=req["end_date"],
                term=12,
                market=None,
                main_sector=req["main_sector"],
                net_rev=req["net_revenue"],
                # net_rev_r=dic_to_list(req["net_revenue_rate"]),
                net_rev_r={'max': None, 'min': None},
                net_prf=req["net_profit"],
                # net_prf_r=dic_to_list(req["net_profit_rate"]),
                net_prf_r={'max': None, 'min': None},
                de_r=req["de_ratio"],
                per=req["per"],
                # psr=dic_to_list(req["psr"]),
                psr={'max': None, 'min': None},
                pbr=req["pbr"],
                # pcr=dic_to_list(req["pcr"]),
                pcr={'max': None, 'min': None},
                op_act=req["activities"]["operating"],
                iv_act=req["activities"]["investing"],
                fn_act=req["activities"]["financing"],
                dv_yld=req["dividend_yield"],
                dv_pay_r=req["dividend_payout_ratio"],
                roa=req["roa"],
                roe=req["roe"],
                marcap=req["market_cap"]
            )

            dt = result["chart"]["start_date"]
            t = dt.timestamp()
            seconds = int(t)
            nanos = int(t % 1 * 1e9)
            proto_timestamp = Timestamp(seconds=seconds, nanos=nanos)
            result["chart"]["start_date"] = proto_timestamp

            return quant_pb2.QuantResult(
                cumulative_return=result["cumulative_return"],
                annual_average_return=result["annual_average_return"],
                winning_percentage=result["winning_percentage"],
                max_loss_rate=result["max_loss_rate"],
                holdings_count=result["holdings_count"],
                chart_data=result["chart"],
            )

        except Exception as e:
            print(e)
            return


def serve():
    print('gRPC server started. Listening to 9000...')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    quant_pb2_grpc.add_QuantServicer_to_server(Quant(), server)

    server.add_insecure_port('[::]:9000')

    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    print('=' * 15 + " Quant Server " + '=' * 15)
    serve()
