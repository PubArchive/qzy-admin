package login

const loginTmpl = `{{define "login_theme1"}}
    <!DOCTYPE html>
    <!--[if lt IE 7]>
    <html class="no-js lt-ie9 lt-ie8 lt-ie7">
    <![endif]-->
    <!--[if IE 7]>
    <html class="no-js lt-ie9 lt-ie8">
    <![endif]-->
    <!--[if IE 8]>
    <html class="no-js lt-ie9">
    <![endif]-->
    <!--[if gt IE 8]><!-->
    <html class="no-js">
    <!--<![endif]-->
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>{{.Title}}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <link rel="stylesheet" href="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.css"}}">
        <style>
          .main-header .logo {
            float: none !important;
            width: 100% !important;
          }
        </style>

        <!--[if lt IE 9]>
        <script src="{{link .CdnUrl .UrlPrefix "/assets/login/dist/respond.min.js"}}"></script>
        <![endif]-->

    </head>
    <body style="margin: 0; background-image: url('{{.BackgroundImage}}'); background-size: cover; background-repeat: no-repeat; background-position: center;">

    <div class="container">
        <div class="row" style="margin-top: 80px;">
            <div class="col-md-4 col-md-offset-4">
                <form action="##" onsubmit="return false" method="post" id="sign-up-form" class="fh5co-form animate-box"
                      data-animate-effect="fadeIn">
                    <h2>{{.Title}}</h2>
                    <div class="form-group">
                        <label for="username" class="sr-only">Username</label>
                        <input type="text" class="form-control" id="username" placeholder="{{lang "username"}}"
                               autocomplete="off">
                    </div>
                    <div class="form-group">
                        <label for="password" class="sr-only">Password</label>
                        <input type="password" class="form-control" id="password" placeholder="{{lang "password"}}"
                               autocomplete="off">
                    </div>
                    <div class="form-group">
                        <button class="btn btn-primary" onclick="submitData()">{{lang "login"}}</button>
                    </div>
                </form>
            </div>
        </div>
    </div>

    <script src="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.js"}}"></script>

    <script>
        function submitData() {
            $.ajax({
                dataType: 'json',
                type: 'POST',
                url: '{{.UrlPrefix}}/signin',
                async: 'true',
                data: {
                    'username': $("#username").val(),
                    'password': $("#password").val()
                },
                success: function (data) {
                    location.href = data.data.url
                },
                error: function (data) {
                    alert('{{lang "login fail"}}');
                }
            });
        }
    </script>

    </body>
    </html>
{{end}}`
