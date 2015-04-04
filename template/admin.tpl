<!DOCTYPE html>
<html lang="zh">
  <head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta charset="utf-8">
    <title>漫画——comic.mozillazg.com</title>
    <meta name="description" content="欢乐氰化物,Cyanide & Happiness 汉化,cyanide and happiness,漫画,comic">
    <meta name="author" content="mozillazg">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="">
    <link rel="shortcut icon" href="">
    <link href="http://libs.baidu.com/bootstrap/3.0.3/css/bootstrap.min.css" rel="stylesheet">
  </head>
  <body>
    <div class="container">
      <div class="main row">
        <h1>admin</h1>
        <p><a href="javascript: void(0);" class="new-comic label label-primary">New</a></p>
        <!-- table -->
        <table class="table table-striped table-bordered table-hover">
          <thead>
              <tr>
                  <th>title</th>
                  <th>url</th>
                  <th>description</th>
                  <th>date</th>
                  <th>action</th>
              </tr>
          </thead>
          <tbody>
              {{range $c := .Comics}}
              <tr>
                  <td>{{$c.Title}}</td>
                  <td>{{$c.ImageURL}}</td>
                  <td>{{$c.Description}}</td>
                  <td>{{$c.Date}}</td>
                  <td>
                    <a href="javascript: void(0);" data-id="{{$c.ID}}"
                      class="edit-comic">edit</a>
                      <a href="javascript: void(0);" data-id="{{$c.ID}}"
                        class="delete-comic">delete</a>
                  </td>
              </tr>
              {{end}}
          </tbody>
        </table>
        <!-- table end -->

      </div>
    </div>

    <!-- #myModal -->
    <div class="modal fade" id="myModal" tabindex="-1" role="dialog"
      aria-hidden="true" >
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title" id="myModalLabel">编辑/新增</h4>
          </div>

          <div class="modal-body">
            <div>
              <form action="#" id="form">
                  <div class="form-group">
                    <label for="title">Title</label>
                    <input class="form-control" type="text" name="title" id="title" required>
                  </div>
                  <div class="form-group">
                    <label for="url">Image URL</label>
                    <input class="form-control" type="text" name="url" id="url" required>
                  </div>
                  <div class="form-group">
                    <label for="description">Description</label>
                    <textarea class="form-control" rows="3" name="description" id="description" required></textarea>
                  </div>
                  <div class="form-group">
                    <label for="date">Date</label>
                    <input class="form-control" name="date" id="date" type="date" date-format="yyyy-mm-dd" required>
                  </div>
            </form>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
              <a class="btn btn-primary" id="post-form" data-id="0"
                href="javascript: void(0);">保存</a>
            </div>

          </div>
        </div>
      </div>
    </div>
    <!-- #myModal end -->


    <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    <script src="http://libs.baidu.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
    <script>
      $(function () {
        $(".new-comic").on("click", function() {
          $("#post-form").data("id");
          $("#myModal").modal("show");
        });
        $(".edit-comic").on("click", function() {
          $("#post-form").data("id", $(this).data("id"));
          $("#myModal").modal("show");
        });

      // new comic
      $("#post-form").on("click", function() {
        var data = $("#form").serialize();
        var id = $(this).data("id");
        var url = "/api/comics";
        var method = "POST";
        if (id != "0") {
          url += "/" + id;
          method = "PUT";
        }
        $.ajax({
          type: method,
          url: url,
          data: data,
          success: function(data) {
            location.reload();
          }
        });
      });
      // delete
      $(".delete-comic").on("click", function() {
        if (!confirm("确定要删除?")) {
          return false;
        }
        var id = $(this).data("id");
        $.ajax({
          type: "DELETE",
          url: "/api/comics/" + id,
          success: function(data) {
            location.reload();
          }
        });
      });

      });
    </script>
  </body>
</html>
