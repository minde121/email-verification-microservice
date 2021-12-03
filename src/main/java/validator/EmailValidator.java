package validator;

import java.util.HashMap;
import java.util.Queue;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class EmailValidator {

    public HashMap<String, Boolean> validate(Queue<String> emails) {

        HashMap<String,Boolean> emailVal = new HashMap<>();

        for (String email : emails) {

            emailVal.put(email, isEmailRegex(email));
        }

        return emailVal;
    }

    private Boolean isEmailRegex(String email) {

        String regex = "^(.+)@(.+)$";

        Pattern pattern = Pattern.compile(regex);
        Matcher matcher = pattern.matcher(email);

        return matcher.matches();
    }
}
