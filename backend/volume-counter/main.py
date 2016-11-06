import schedule, time, sys, db, counter
from config import timimg_config

def main():
    try:
        tickers = _get_unique_tickers()
        _run_service(tickers)
    except KeyboardInterrupt as e:
        db.close_connection()
        print " Exiting"
        sys.exit(0)

def _run_service(tickers):
    #schedule.every().minute.do(counter.count_volume, tickers)
    while True:
        #schedule.run_pending()
        counter.count_volume(tickers)
        time.sleep(timimg_config["SLEEP"])

def _get_unique_tickers():
    db_result = db.get_unique_tickers()
    if db_result["success"]:
        return map(lambda row: counter.format_ticker(row[0]), db_result["data"])
    else:
        sys.exit(1)

if __name__ == '__main__':
    main()
