{{define `main`}}
<div id="main" class="container">
    <div class="row" id="ct">
        <div class="col s12">
            <p></p>
            <h4>Your feedback</h4>
        </div>
        <div class="col s12">
            <p>
               We' appreciate your feedback on this app and the code generators for
                embedded code! Please let us know what you think and how we could improve
                this service. Additionally, feel free to use this form to contact us.
            </p>
            <p>
                We're fans of Datensparsamkeit/data minimization: Please provide personal
                contact information only if you agree that we will process it for the purpose of contacting you.
            </p>
            <br/>
            <form id="fbf" name="fbf" method="POST" action="/feedback">
                <div class="row">
                    <div class="input-field col s12">
                        <select id="fbf_what" name="fbf_what">
                            <option value="none" selected>Please choose a category</option>
                            <option value="ui">Feedback on UI/App</option>
                            <option value="gen">Feedback on Generators</option>
                            <option value="support">Support for a board/framework/...</option>
                            <option value="other">Other...</option>
                        </select>
                        <label>Select a type.</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <textarea placeholder="Your feedback as text" class="materialize-textarea" name="fbf_text" id="fbf_text" size="20">{{ .CtfDesc }}</textarea>
                        <label for="ctf_desc">Text</label>
                    </div>
                </div>
                <div class="row" >
                    <a class="waves-effect waves-light deep-orange btn right" id="fbf_send"><i class="material-icons right">send</i>Send</a>
                </div>
            </form>
        </div>
    </div>
</div>
{{end}}