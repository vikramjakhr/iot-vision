<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <link rel="apple-touch-icon" sizes="76x76" href="/static/img/apple-icon.png">
    <link rel="icon" type="image/png" sizes="96x96" href="/static/img/favicon.ico">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>

    <title>Text & Face Detection</title>

    <meta content='width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0' name='viewport'/>
    <meta name="viewport" content="width=device-width"/>


    <!-- Bootstrap core CSS     -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet"/>

    <!-- Animation library for notifications   -->
    <link href="/static/css/animate.min.css" rel="stylesheet"/>

    <!--  Paper Dashboard core CSS    -->
    <link href="/static/css/paper-dashboard.css" rel="stylesheet"/>


    <!--  CSS for Demo Purpose, don't include it in your project     -->
    <link href="/static/css/demo.css" rel="stylesheet"/>


    <!--  Fonts and icons     -->
    <link href="http://maxcdn.bootstrapcdn.com/font-awesome/latest/css/font-awesome.min.css" rel="stylesheet">
    <link href='https://fonts.googleapis.com/css?family=Muli:400,300' rel='stylesheet' type='text/css'>
    <link href="/static/css/themify-icons.css" rel="stylesheet">

</head>
<body>

<div class="wrapper">
    <div class="sidebar" data-background-color="white" data-active-color="danger">

        <!--
            Tip 1: you can change the color of the sidebar's background using: data-background-color="white | black"
            Tip 2: you can change the color of the active button using the data-active-color="primary | info | success | warning | danger"
        -->

        <div class="sidebar-wrapper">
            <div class="logo">
                <a href="http://www.tothenew.com" class="simple-text">
                    TTN
                </a>
            </div>

            <ul class="nav">
                <li class="active">
                    <a href="/">
                        <i class="ti-panel"></i>
                        <p>Text & Face Detection</p>
                    </a>
                </li>
            </ul>
        </div>
    </div>

    <div class="main-panel">
        <nav class="navbar navbar-default">
            <div class="container-fluid">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar bar1"></span>
                        <span class="icon-bar bar2"></span>
                        <span class="icon-bar bar3"></span>
                    </button>
                    <a class="navbar-brand" href="#">Dashboard</a>
                </div>
            </div>
        </nav>


        <div class="content">
            <div id="data" class="container-fluid" style="width: 70%">
                <!--<div class="row">-->
                <!--<div class="col-lg-12 col-md-12">
                    <div class="card card-user" style="height: auto !important;">
                        <div class="image" style="height: auto !important;">
                            <img src="https://qph.ec.quoracdn.net/main-qimg-f93403f6d32bc43b40d85bd978e88bbf"
                                 alt="..."/>
                        </div>
                    </div>
                </div>
                <div class="col-lg-6 col-md-6">
                    <div class="card card-user" style="height: auto !important;">
                        <div class="content">
                            <p class="description text-center">
                                "I like the way you work it <br>
                                No diggity <br>
                                I wanna bag it up"
                            </p>
                        </div>
                    </div>
                </div>
                <div class="col-lg-6 col-md-6">
                    <div class="card card-user" style="height: auto !important;">
                        <div class="image" style="height: auto !important;">
                            <img src="https://qph.ec.quoracdn.net/main-qimg-f93403f6d32bc43b40d85bd978e88bbf"
                                 alt="..."/>
                        </div>
                    </div>
                </div>-->
                <!--</div>-->
            </div>
        </div>


        <footer class="footer">
            <div class="container-fluid">
                <nav class="pull-left">
                    <ul>

                        <li>
                            <a href="http://www.tothenew.com">
                                TTN
                            </a>
                        </li>
                    </ul>
                </nav>
                <div class="copyright pull-right">
                    &copy;
                    <script>document.write(new Date().getFullYear())</script>
                    , made with <i class="fa fa-heart heart"></i> by <a href="http://www.tothenew.com">TTN</a>
                </div>
            </div>
        </footer>

    </div>
</div>


</body>

<!--   Core JS Files   -->
<script src="/static/js/jquery-1.10.2.js" type="text/javascript"></script>
<script src="/static/js/bootstrap.min.js" type="text/javascript"></script>

<!--  Checkbox, Radio & Switch Plugins -->
<script src="/static/js/bootstrap-checkbox-radio.js"></script>

<!--  Charts Plugin -->
<script src="/static/js/chartist.min.js"></script>

<!--  Notifications Plugin    -->
<script src="/static/js/bootstrap-notify.js"></script>

<!-- Paper Dashboard Core javascript and methods for Demo purpose -->
<script src="/static/js/paper-dashboard.js"></script>

<script type="text/javascript">
    setInterval(function () {
        $.ajax({
            url: "/textReco",
            type: 'GET',
            error: function () {
                console.log("error")
            },
            success: function (data, status) {
                if (data.Url !== "") {
                    var html = "";
                    html += '<div class="row"> <div class="col-lg-12 col-md-12"> \
                        <div class="card card-user" style="height: auto !important;"> \
                        <div class="image" style="height: auto !important;"> \
                        <img src="' + data.Url + '" alt="..."/> \
                        </div> \
                        </div> \
                        </div> \
                        <div class="col-lg-6 col-md-6"> \
                        <div class="card card-user" style="height: auto !important;"> \
                        <div class="content"> \
                        <p class="description text-center">' + data.Text +

                        '</p> \
                        </div> \
                        </div> \
                        </div>';
                    if (data.FaceInfoList.length > 0) {
                        data.FaceInfoList.forEach(function (item) {
                            html += '<div class="col-lg-6 col-md-6"> \
                            <div class="card card-user" style="height: auto !important;"> \
                            <div class="image" style="height: auto !important;"> \
                        <img src="' + item.Url + '" alt="..."/> \
                        </div> \
                        <div class="content" style="min-height: auto"> \
                        <p class="description"><b>Confidence:&nbsp;&nbsp;</b>' + item.Confidence + '</p> \
                        <p class="description"><b>Similarity:&nbsp;&nbsp; </b>' + item.Similarity + '</p> \
                        </div> \
                        </div> \
                        </div>'
                        });
                    } else {
                        html += '<div class="col-lg-6 col-md-6"> \
                    <div class="card card-user" style="height: auto !important;"> \
                    <div class="content"> \
                    <p class="description text-center">' + "No Image matched" +
                            '</p> \
                            </div> \
                            </div> \
                            </div>';
                    }
                    html += "</div>"

                }
                $("#data").append(html);
            }
        });
    }, 1000);
</script>


</html>
