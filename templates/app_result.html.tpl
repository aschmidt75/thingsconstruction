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
                    <p>Finally, code is ready! You can download files individually from the table below,
                    or download all as an archive.</p>
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
                                <td>WoT Thing Description</td>
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
                                    <a href="/app/{{$GThingId}}/result/asset/{{ .Permalink }}"><i class="material-icons">file_download</i>
                                    </a>
                                    <a target="#" class="modal-trigger" href="#view_modal"> <i id="view_{{ .Permalink }}" linkid="{{ .Permalink }}" class="material-icons">remove_red_eye</i>
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
                <a href="#!" class="modal-action modal-close waves-effect btn-flat">Dismiss</a>
            </div>
        </div>
{{end}}