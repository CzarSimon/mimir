import db, math, time, sys, json, req
import numpy as np
from datetime import datetime, date
from collections import Counter
from config import reciving_server

def retrive_and_calc(tickers):
    for ticker in tickers:
        db_result = db.get_tweets_for_ticker(ticker)
        if db_result["success"]:
            timestamps = _format_timestamps(db_result["data"])
            result = _filter_and_calc(timestamps, ticker)
            result["ticker"] = ticker.replace("$", "")
            _send_result(result)
        else:
            print db_result["data"]

def _format_timestamps(timestamps):
    return map(lambda row: datetime.strftime(row[0], "%Y-%m-%d:%H"), timestamps)

def _filter_and_calc(timestamps, ticker):
    result = {}
    filtered_timestamps = _separate_busdays(timestamps)
    days = _calc_days_meassured(ticker)
    for key, timestamps in filtered_timestamps.iteritems():
        hourly_volumes = _reduce_by_day(_reduce_by_hour(timestamps), days[key])
        mean, stdev = _calc_mean_stdev(hourly_volumes)
        result[key] = {"mean": mean, "stdev": stdev}
    return result

def _calc_mean_stdev(hourly_volumes):
    mean_list = map(lambda volumes: round(np.mean(volumes), 2), hourly_volumes)
    stdev_list = map(lambda volumes: round(np.std(volumes), 2), hourly_volumes)
    return mean_list, stdev_list

def _reduce_by_day(volume_day_hour, no_days):
    hourly_volumes = [0] * 24
    for hour in range(0, 24):
        hour_str = str(hour) if hour > 9 else "0" + str(hour)
        temp = filter(lambda item: item[0].endswith(hour_str), volume_day_hour.iteritems())
        hourly_volumes[hour] = _add_missing_days(map(lambda item: float(item[1]), temp), no_days)
    return hourly_volumes

def _add_missing_days(volume_list, no_days):
    list_length = len(volume_list)
    if list_length < no_days:
        return volume_list + [0.0] * (no_days - list_length)
    else:
        return volume_list

def _separate_busdays(timestamps):
    busdays = filter(lambda date_str: _is_busday(date_str), timestamps)
    weekdays = filter(lambda date_str: not _is_busday(date_str), timestamps)
    return {"busdays": busdays, "weekend_days": weekdays}

def _reduce_by_hour(timestamps):
    return dict(Counter(timestamps))

def _is_busday(date_str):
    dates = [date_str.split(":")[0]]
    return np.is_busday(dates)[0]

def _calc_days_meassured(ticker):
    db_result = db.get_first_day_stored(ticker)
    if db_result["success"]:
        start_date = datetime_to_date(db_result["data"][0][0])
        today = datetime_to_date(datetime.now())
        bus_days = int(np.busday_count(start_date, today)) + 1
        all_days = (today - start_date).days + 1
        return {"busdays": bus_days, "weekend_days": (all_days - bus_days)}

def datetime_to_date(dt):
    return date(dt.year, dt.month, dt.day)

def _send_result(result):
    print result
    endpoint = "".join([reciving_server["ADDRESS"], reciving_server["ROUTE"]])
    request_data = json.dumps({"data": result, "message": "Updated mean and stdev for {}".format(result["ticker"])})
    req.send_stats_to_server(endpoint, request_data, result["ticker"])
