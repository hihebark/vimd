package core

//TEMPLATE for the front-end.
const TEMPLATE = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>{{.Current}}</title>
  <link crossorigin="anonymous" media="all" integrity="sha512-RPWwIpqyjxv5EpuWKUKyeZeWz9QEzIbAWTiYOuxGieUq7+AMiZbsLeQMfEdyEIUoNjLagHK0BEm92BmXnvaH4Q==" rel="stylesheet" href="https://github.githubassets.com/assets/frameworks-40c1c9d8ff06284fb441108e6559f019.css" />
  <link crossorigin="anonymous" media="all" integrity="sha512-tCm5Qwdem1lB66O3j+wtOkAZAZ2vqNmh9sMLzmRhjZzW2jn3uAJSyt3x4p52+fz6b0MTf/VUjaXxzsgsvN8HSg==" rel="stylesheet" href="https://github.githubassets.com/assets/github-aeab953696d5337e9465738f175ce344.css" />
  <style>
    /* Page tweaks */
    .preview-page {
      margin-top: 20px;
    }
    /* User-content tweaks */
    .timeline-comment-wrapper > .timeline-comment:after,
    .timeline-comment-wrapper > .timeline-comment:before {
      content: none;
    }
    /* User-content overrides */
    .discussion-timeline.wide {
      width: 920px;
    }
  </style>
</head>
<body>
  <div class="page">
    <div class="file-wrap container">
      <table class="files js-navigation-container js-active-navigation-container" data-pjax>
        <tbody>
          {{ range $k, $v := .Files }}
          <tr class="js-navigation-item">
            <td class="icon">
              <svg class="octicon octicon-file" viewBox="0 0 12 16" version="1.1" width="12" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M6 5H2V4h4v1zM2 8h7V7H2v1zm0 2h7V9H2v1zm0 2h7v-1H2v1zm10-7.5V14c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V2c0-.55.45-1 1-1h7.5L12 4.5zM11 5L8 2H1v12h10V5z"/></svg>
              <img width="16" height="16" class="spinner" alt="" src="https://assets-cdn.github.com/images/spinners/octocat-spinner-32.gif" />
            </td>
            <td class="content">
              <span class="css-truncate css-truncate-target">
                <a class="js-navigation-open" title="util.go" id="" href="/{{ $v.Name }}">{{ $v.Name }}</a>
              </span>
            </td>
            <td class="message">{{ $v.Metadata.Commit }}</td>
            <td class="age">{{$v.Metadata.Date}}</td>
          </tr>
          {{ end }}
        </tbody>
      </table>
    </div>
    <div id="preview-page" class="preview-page" data-autorefresh-url="">
      <div role="main" class="main-content">
        <div class="container new-discussion-timeline experiment-repo-nav">
          <div class="repository-content">
            <div id="readme" class="readme boxed-group clearfix announce instapaper_body md">
              <h3>
				<svg class="octicon octicon-book" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M3 5h4v1H3V5zm0 3h4V7H3v1zm0 2h4V9H3v1zm11-5h-4v1h4V5zm0 2h-4v1h4V7zm0 2h-4v1h4V9zm2-6v9c0 .55-.45 1-1 1H9.5l-1 1-1-1H2c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1h5.5l1 1 1-1H15c.55 0 1 .45 1 1zm-8 .5L7.5 3H2v9h6V3.5zm7-.5H9.5l-.5.5V12h6V3z"/></svg>
				<span class="octicon octicon-book"></span>
				{{.Current}}
			  </h3>
              <article class="markdown-body entry-content" itemprop="text" id="grip-content">
                {{.Content}}
              </article>
            </div>
          </div>
        </div>
      </div>
  </div>
  <p align="center"><b>Vis-md</b> - <i>Made by <b><a href="https://github.com/hihebark">hihebark</a></b></i></p>
</body>
</html>`
const NOTFOUNDPAGE = `<h1>NOT FOUND PAGE</h1>`
