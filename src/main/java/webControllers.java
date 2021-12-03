import com.google.gson.Gson;
import ds.EmailCollector;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import validator.EmailValidator;

import java.util.HashMap;

@Controller
public class webControllers {

    EmailValidator emailValidator = new EmailValidator();

    @RequestMapping(value = "/email/validateEmails", method = RequestMethod.POST)
    @ResponseStatus(value = HttpStatus.OK)
    @ResponseBody
    public String validateEmails(@RequestBody String request){

        Gson gson = new Gson();

        EmailCollector emails = gson.fromJson(request, EmailCollector.class);

        HashMap<String, Boolean> emailValidation = emailValidator.validate(emails.getEmails());

        return gson.toJson(emailValidation);
    }

}
