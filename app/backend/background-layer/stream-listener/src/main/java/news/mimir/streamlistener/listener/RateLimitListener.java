package news.mimir.streamlistener.listener;

import twitter4j.RateLimitStatus;
import twitter4j.RateLimitStatusEvent;
import twitter4j.RateLimitStatusListener;

import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;

/**
 * @author simon.g.lindgren@gmail.com
 */
public class RateLimitListener implements RateLimitStatusListener {

    private static final Logger log =
            Logger.getLogger(RateLimitListener.class.getName());

    @Override
    public void onRateLimitStatus(RateLimitStatusEvent rateLimitStatusEvent) {
        log.info(rateLimitStatusEvent.getRateLimitStatus().toString());
    }

    @Override
    public void onRateLimitReached(RateLimitStatusEvent rateLimitStatusEvent) {
        RateLimitStatus rateLimit = rateLimitStatusEvent.getRateLimitStatus();
        log.warning(getRateLimitMessage(rateLimit));
        try {
            TimeUnit.SECONDS.sleep(rateLimit.getResetTimeInSeconds());
        } catch (InterruptedException e) {
            log.severe(e.getMessage());
        }
    }

    /**
     * Creates rate limit message
     * @param rateLimit
     * @return
     */
    private String getRateLimitMessage(RateLimitStatus rateLimit) {
        return String.format("Rate limit reached, sleeping for %s seconds", rateLimit.getResetTimeInSeconds());
    }

}
