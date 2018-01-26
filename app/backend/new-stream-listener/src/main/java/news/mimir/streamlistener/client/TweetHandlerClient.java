package news.mimir.streamlistener.client;

import com.mashape.unirest.http.HttpResponse;
import com.mashape.unirest.http.JsonNode;
import com.mashape.unirest.http.Unirest;
import com.mashape.unirest.http.async.Callback;
import com.mashape.unirest.http.exceptions.UnirestException;
import news.mimir.streamlistener.config.Config;
import news.mimir.streamlistener.config.ServerConfig;
import org.json.JSONArray;

import java.util.concurrent.Future;
import java.util.logging.Logger;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class TweetHandlerClient implements HttpClient {

    private static final Logger log =
            Logger.getLogger(TweetHandlerClient.class.getName());

    private final int STATUS_OK = 200;

    private ServerConfig config;
    private Callback callback = setupCallback();

    /**
     * Constructor
     */
    public TweetHandlerClient() {
        this.config = Config.SERVER_CONFIG;
    }

    /**
     * (non-Javadoc)
     * @see HttpClient#send(String)
     */
    @Override
    public void send(String message) {
        log.info("Forwarding tweet");
        makePostRequest(config.POST_TWEET_ROUTE, message);
    }

    /**
     * (non-Javadoc)
     * @see HttpClient#getFilter()
     */
    @Override
    public String[] getFilter() {
        HttpResponse<JsonNode> response = makeJsonGetRequest(config.GET_TICKERS_ROUTE);
        return formatTickers(response.getBody().getArray());
    }

    private void makePostRequest(String URL, String body) {
        Unirest.post(URL)
                .header("Content-Type", "application/json")
                .body(body)
                .asStringAsync(callback);
    }

    private Callback<String> setupCallback() {
        return new Callback<String>() {
            @Override
            public void completed(HttpResponse<String> httpResponse) {
                if (httpResponse.getStatus() == STATUS_OK) {
                    log.info("Tweet sent and accepted");
                } else {
                    log.warning(String.format(
                            "%s - %s", httpResponse.getStatus(), httpResponse.getBody()));
                }
            }

            @Override
            public void failed(UnirestException e) {
                log.warning(e.getMessage());
            }

            @Override
            public void cancelled() {
                log.warning("Request cancelled");
            }
        };
    }

    /**
     * Formats tickers to track.
     * @param tickers
     * @return
     */
    private String[] formatTickers(JSONArray tickers) {
        String[] formatedTickers = new String[tickers.length()];
        for (int i = 0; i < tickers.length(); i++) {
            formatedTickers[i] = "$" + tickers.getString(i);
        }
        return formatedTickers;
    }

    /**
     * Makes a get get request to a supplied URL, throws a runtime exception in case of failures.
     * @param URL
     * @throws RuntimeException
     * @return response
     */
    private HttpResponse<JsonNode> makeJsonGetRequest(String URL) {
        try {
            HttpResponse<JsonNode> response = Unirest.get(URL).asJson();
            if (response.getStatus() != STATUS_OK) {
                throw new UnirestException("Non 200 response, Status: %d" + response.getStatus());
            }
            return response;
        } catch (UnirestException e) {
            log.severe(e.getMessage());
            throw new RuntimeException(e.getMessage());
        }
    }
}
