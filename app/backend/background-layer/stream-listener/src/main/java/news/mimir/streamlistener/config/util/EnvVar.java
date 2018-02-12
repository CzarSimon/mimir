package news.mimir.streamlistener.config.util;

/**
 * @author simon.g.lindgren@gmail.com
 */
public final class EnvVar {

    /**
     * Private constructor to prevent instantiation
     */
    private EnvVar() {}

    /**
     * Attempts to get an environment variable based on its name
     * @param key
     * @return value of environment variable
     */
    public static String get(String key) {
        String value = System.getenv(key);
        if (value == null) {
            throw new EnvNotFoundException(key);
        }
        return value;
    }

    /**
     *
     * @param key
     * @param defaultValue
     * @return
     */
    public static String get(String key, String defaultValue) {
        String value = System.getenv(key);
        if (value == null) {
            return defaultValue;
        }
        return value;
    }

}
