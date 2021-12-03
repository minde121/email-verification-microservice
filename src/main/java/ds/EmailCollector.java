package ds;

import java.util.Queue;

public class EmailCollector {

    private Queue<String> emails;

    public EmailCollector(Queue<String> emails) {
        this.emails = emails;
    }

    public EmailCollector() {
    }

    public Queue<String> getEmails() {
        return emails;
    }

    public void setEmails(Queue<String> emails) {
        this.emails = emails;
    }

}
