import schedule, time, sys, db, counter
from config import timimg

def main():
    print "Running volume counter"
    try:
        tickers = _get_unique_tickers()
        print tickers, "\n"
        _run_service(tickers)
    except KeyboardInterrupt as e:
        db.close_connection()
        print " Exiting"
        sys.exit(0)

def _run_service(tickers):
    counter.count_volume(tickers)
    schedule.every().minute.do(counter.count_volume, tickers)
    while True:
        schedule.run_pending()
        time.sleep(timimg["SLEEP"])

def _get_unique_tickers():
    db_result = db.get_unique_tickers()
    if db_result["success"]:
        return map(lambda row: counter.format_ticker(row[0]), db_result["data"])
    else:
        sys.exit(1)

if __name__ == '__main__':
    main()
