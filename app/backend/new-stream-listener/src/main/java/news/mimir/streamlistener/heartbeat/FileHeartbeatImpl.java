package news.mimir.streamlistener.heartbeat;

import news.mimir.streamlistener.heartbeat.config.Config;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.logging.Logger;

/**
 * Implementation of file heartbeat emitter.
 * @author simon.g.lindgren@gmail.com
 */
public class FileHeartbeatImpl implements FileHeartbeat {

    private static final Logger log =
            Logger.getLogger(FileHeartbeatImpl.class.getName());

    private String filepath;
    private long interval;

    private static final int MS_IN_SECOND = 1000;

    /**
     * Constructor
     */
    public FileHeartbeatImpl() {
        Config config = new Config();
        this.filepath = config.HEARTBEAT_FILE;
        this.interval = MS_IN_SECOND * config.HEARTBEAT_INTERVAL_SECONDS;
    }

    @Override
    public void runHeartbeat() {
        while (!Thread.interrupted()) {
            emitToFile();
            try {
                Thread.sleep(interval);
            } catch (InterruptedException e) {
                break;
            }
        }
    }

    @Override
    public void emitToFile() {
        File file = new File(filepath);
        try {
            touch(file, System.currentTimeMillis());
        } catch (IOException e) {
            log.severe(e.getMessage());
            throw new RuntimeException(e.getMessage());
        }
    }

    /**
     *
     * @param file
     * @param timestamp
     * @throws IOException
     */
    private void touch(File file, long timestamp) throws IOException {
        if (!file.exists()) {
            new FileOutputStream(file).close();
        }
        file.setLastModified(timestamp);
    }
}
