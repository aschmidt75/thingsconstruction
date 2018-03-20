{{define `main`}}
        {{ $GThingId := .ThingId }}
<div id="main" class="no-container">
    <div class="row">
        <div class="col s12">
        {{ if .Message }}
            <div class="card red darken-2">
                <div class="card-content white-text">
                {{ .Message }}
                </div>
            </div>
        {{ else }}
            <div class="row" id="gen">
                <div class="col s9">
                    <h4>Download your code!</h4>
                    <p>Finally, code is ready! You can view and download files individually from the table below,
                    or download all as an archive.
                    <br>
                    <span class="tc-maincolor-text" id="btn_more">... more</span>
                    <span id="sp_more" class="hide">
                        The table below shows all generated assets, which include the Thing/API description,
                        source code, readme's / how-to's and licensing information. The rightmost buttons
                        allow you to download the file(s) to your pc, or view them in the browser. From browser
                        view, code can be copy&pasted, so it's easily transferable to an IDE for example.
                        A good start is the README file which contains detailed instructions
                        of how to get the code running on your device.
                        <span class="tc-maincolor-text" id="btn_less">... less</span>
                    </span>

                    </p>
                    <p>You can revise your work later by coming back to this page:<br>
                        <i data-tooltip="Copy to clipboard" id="url_copy_btn" class="material-icons tooltipped">content_copy</i>
                        <input type="text" id="url_copy_input" class="hide">
                        <a id="url_copy_a" href="{{.URLPrefix}}app/{{.ThingId}}/generate">{{.URLPrefix}}app/{{.ThingId}}/generate</a>
                    </p>
                    <div>
                        <table>
                            <thead>
                                <tr>
                                    <th>Filename</th>
                                    <th>Description</th>
                                    <th>Actions</th>
                                </tr>
                            </thead>
                            <tr>
                                <td>{{ .ThingId }}.json</td>
                                <td>Thing/API Description</td>
                                <td>
                                    <a href="/app/{{.ThingId}}/result/wtd">
                                    <i class="material-icons">file_download</i>
                                </a>
                                </td>
                            </tr>
                            {{ if .Files }}
                                {{ range .Files }}
                            <tr>
                                <td>{{ .FileName }}</td>
                                <td>{{ .Description }}<br>
                                {{ .FileType }} / {{ .Language }}</td>
                                <td>
                                    <a href="/app/{{$GThingId}}/result/asset/{{ .Permalink }}"><i class="material-icons tooltipped" data-delay="100" data-tooltip="Download file">file_download</i>
                                    </a>
                                    <a target="#" class="modal-trigger" href="#view_modal"> <i id="view_{{ .Permalink }}" linkid="{{ .Permalink }}" onclick="return view_element(event);" class="material-icons tooltipped" data-delay="100" data-tooltip="View in browser">remove_red_eye</i>
                                    </a>
                                </td>
                            </tr>
                                {{ end }}
                            {{ end }}
                        </table>
                    </div>
                    <div style="padding-top:2em">

                        <div class="input-field col s6" >
                            <select id="select-archive">
                                <option value="" disabled selected>your favorite archive format</option>
                                <option value="zip">ZIP Archive (.zip)</option>
                                <option value="tar">Tarball (.tar)</option>
                                <option value="targz">GZipped Tarball (.tar.gz)</option>
                            </select>
                            <label>Download all as ...</label>
                        </div>
                        <div class="input-field col s2" >
                            <a href="/app/{{$GThingId}}/result/asset-archive/" class="hide" id="select-archive-link"><i class="material-icons">file_download</i></a>
                            </div>
                    </div>
                </div>
                <div class="col s2 offset-s1 hide-on-med-and-down">
                    <br/>
                    <p>
                    Need to change some things? Jump back to individual
                    Thing Description parts here:
                    <ul class="section table-of-contents">
                        <li><a href="/app/{{.ThingId}}">Create</a></li>
                        <li><a href="/app/{{.ThingId}}/framework">Framework</a></li>
                        <li><a href="/app/{{.ThingId}}/properties">Properties</a></li>
                        <li><a href="/app/{{.ThingId}}/actions">Actions</a></li>
                        <li><a href="/app/{{.ThingId}}/events">Events</a></li>
                        <li><a href="/app/{{.ThingId}}/generate">Generate</a></li>
                    </ul>
                    </p>
                    {{ if .Feature.Flattr }}
                    <p>
                        <a target="tcext" href="https://flattr.com/submit/auto?user_id={{ .FlattrUser }}&url=https://thngstruction.online/&title=&language=en&tags=app&category=software">
                            <img src="https://button.flattr.com/flattr-badge-large.png">
                        </a>
                    </p>
                    {{ end }}
                </div>

            </div>

        {{ end }}
        </div>
    </div>
</div>
        <!-- Modal Structure -->
        <div id="view_modal" class="modal modal-fixed-footer">
            <div class="modal-content" id="view_modal_content">
            </div>
            <div class="modal-footer">
                <a id="copycode" data-clipboard-target="#view_modal_content" class="modal-action waves-effect tc-maincolor btn-flat">Copy to clipboard</a>
                <a class="modal-action modal-close waves-effect btn-flat">Dismiss</a>
            </div>
        </div>

        <span id="span_thing_id" class="hide">{{.ThingId}}</span>

{{end}}