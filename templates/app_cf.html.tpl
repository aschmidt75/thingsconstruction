{{define `main`}}
<div id="main" class="no-container">
    <div class="row">
        <div class="col s12">
            <!-- cf == choose framework -->
            <div class="col s9">
                <div class="row">
                    <p><br></p>
                    <h4>Choose Code Generation Module</h4>
                    <p>
                        The list below shows all available generators. You can narrow down the list of
                        generators by selecting implementation aspects from the stack below.
                        <br>
                        <span class="tc-maincolor-text" id="btn_more">... more</span>
                        <span id="sp_more" class="hide">
                            A generator is needed to turn your device model into runnable code for a
                            specific embedded devices. The generator module determines where code can run
                            according to an embedded platform, what protocols it uses etc. By default,
                            the list shows all available generators. by selecting implementation items
                            from the chips below you can narrow down the list.<br>
                            Clicking on a generator shows additional information about it, clicking on
                            <i class="material-icons tiny">description</i> opens a detailed description
                            in a separate browser tab. Clicking on <i class="material-icons tiny">arrow_forward</i>
                            selects the module for the upcoming steps.
                            <span class="tc-maincolor-text" id="btn_less">... less</span>
                        </span>
                    </p>
                    <div class="row">
                        <div class="col s2 tc-maincolor-text darken-1">
                            Available
                        </div>
                        <div id="cf_targets_available" class="col s8 teal-text">
                        </div>
                    </div>
                    <div class="row">
                        <div class="col s2 tc-maincolor-text darken-1">
                            Filter by ...
                        </div>
                        <div id="cf_targets_selected" class="col s8 teal-text">
                        </div>
                    </div>

                </div>
                <div class="row">
                    <h4>Available generators<span class="new badge grey darken-1" id="cf_targets_num_showing" data-badge-caption="">0</span></h4>
                    <ul class="collapsible" data-collapsible="accordion" id="cf_targets_matching">
                    </ul>
                    <p></p>
                </div>
                <div class="hide">
                    <form id="cf_selection_form" name="cf_selection_form" method="POST" action="/app/{{ .AppPageData.ThingId }}/framework">
                        <input type="text" id="cf_selection" name="cfs">
                        <input type="text" id="cf_id" name="cfid" value="{{ .AppPageData.ThingId }}">
                    </form>
                </div>
                {{ if .Feature.VoteForGenerators }}
                <div class="row">
                    <span id="span_btn_interest">
                    <a id="btn_interest" class="waves-effect waves-light btn-large tc-maincolor" href="#"><i class="material-icons left">extension</i>Interested in additional generators?</a>
                    </span>
                    <span id="span_interest" class="hide">
                    <form id="cf_vote_form" name="cf_vote_form" method="POST" action="/feedback/vote">
                        <p id="cf_vote_scrollTo">Then please vote here for your favorite IoT tech stack:</p>
                    <table>
                {{ range $k, $v := .VoteGenerators }}
                        <tr>
                            <td>
                                <input type="text" class="hide" id="vote_{{$k}}-input" name="{{$k}}" value="3">
                                <div class="col s2">
                                <i id="vote_{{$k}}-down" class="material-icons">thumb_down&nbsp;</i>
                                </div>
                                <div id="vote_{{$k}}-1" class="col s1 ">&nbsp;</div>
                                <div id="vote_{{$k}}-2" class="col s1 ">&nbsp;</div>
                                <div id="vote_{{$k}}-3" class="col s1 tc-maincolor">&nbsp;</div>
                                <div id="vote_{{$k}}-4" class="col s1 ">&nbsp;</div>
                                <div id="vote_{{$k}}-5" class="col s1 ">&nbsp;</div>
                                <div class="col s2">
                                <i id="vote_{{$k}}-up" class="material-icons">&nbsp;thumb_up</i>
                                </div>
                            </td>
                            <td><strong>{{$v}}</strong></td>
                        </tr>
                {{ end }}
                    </table>
                    <button id="vote_submit" class="btn waves-effect waves-light tc-maincolor" href="#!">Vote!</button>
                    </form>
                    </span>
                </div>
                {{ end }}
            </div>
            <div class="col s2 offset-s1 hide-on-med-and-down">
                <p>
                <ul class="section table-of-contents">
                    <li>Create</li>
                    <li><strong class="tc-maincolor-text">Framework</strong>{{ if .TocInfo.framework }}&nbsp;(<span>{{ index .TocInfo "framework" }}</span>){{end}}</li>
                    <li>Properties{{ if .TocInfo.num_properties }}&nbsp;(<span>{{ index .TocInfo "num_properties" }}</span>){{end}}</li>
                    <li>Actions{{ if .TocInfo.num_events }}&nbsp;(<span id="toc_current_info">{{ index .TocInfo "num_actions" }}</span>){{end}}</li>
                    <li>Events{{ if .TocInfo.num_events }}&nbsp;(<span>{{ index .TocInfo "num_events" }}</span>){{end}}</li>
                    <li>Generate!</li>
                </ul>
                </p>
            </div>
        </div>
    </div>
</div>
{{end}}