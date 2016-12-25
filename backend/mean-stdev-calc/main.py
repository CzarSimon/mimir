import sys, time, schedule, calc, db
from config import service_name, timimg

def main():
    print "Running {}".format(service_name)
    try:
        tickers = _get_unique_tickers()
        _run_service(tickers)
    except KeyboardInterrupt as e:
        db.close_connection()
        print " Exiting"
        sys.exit(0)

def _run_service(tickers):
    calc.retrive_and_calc(tickers) # For testing imidiate result
    schedule.every().day.at(timimg["EXEC_TIME"]).do(calc.retrive_and_calc, tickers)
    while True:
        schedule.run_pending()
        time.sleep(timimg["SLEEP"])

def _get_unique_tickers():
    db_result = db.get_unique_tickers()
    if db_result["success"]:
        return map(lambda row: _format_ticker(row[0]), db_result["data"])
    else:
        sys.exit(1)

def _format_ticker(ticker):
    return ticker
    #return ticker.replace("$", "")

if __name__ == '__main__':
    main()
