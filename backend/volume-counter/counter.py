from datetime import datetime
from config import reciving_server
import db, sys, json, req

def _send_results(volume_count, timestamp):
    endpoint = "".join([reciving_server["ADDRESS"], reciving_server["ROUTE"]])
    print "Current volumes:", volume_count, "end\n"
    req_data = json.dumps({"data": volume_count, "message": "sending latest volume count"})
    req.post_volumes(endpoint, req_data, {'content-type': 'application/json'}, timestamp)

def _get_current_hour(now):
    return datetime.strftime(now, "%Y-%m-%d %H") + ":00:00"

def _retrive_count(current_hour, all_tickers):
    db_result = db.get_hourly_volume(current_hour)
    if db_result["success"]:
        return _format_volume_count(db_result["data"], all_tickers)

def _format_volume_count(retrived, all_tickers):
    retrived_volumes = { format_ticker(row[0]): int(row[1]) for row in retrived }
    return { ticker: (retrived_volumes[ticker] if (ticker in retrived_volumes) else 0) for ticker in all_tickers }

def format_ticker(ticker):
    return ticker.replace("$", "")

def count_volume(tickers):
    now = datetime.utcnow()
    volume_count = _retrive_count(_get_current_hour(now), tickers)
    _send_results(volume_count, now)
