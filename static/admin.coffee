$ ->
  $(".new-comic").on "click", ->
    $("#form")[0].reset()
    $("#post-form").data "id"
    $("#myModal").modal "toggle"

  $(".edit-comic").on "click", =>
    $("#form")[0].reset()
    $("#post-form").data "id", $(_this).data "id"

    $("input#title").val $(_this).data("title").trim()
    $("input#url").val $(_this).data("url").trim()
    $("textarea#description").val $(_this).data("desc").trim()
    $("input#date").val $(_this).data("date").trim()

    $("#myModal").modal "toggle"

  # new comic
  $("#post-form").on "click", =>
    data = $("#form").serialize()
    id = $(_this).data("id")
    url = "/api/comics"
    method = "POST"
    if id isnt "0"
      url = "#{ url }/#{ id }"
      method = "PUT"

    $.ajax
      type: method
      url: url
      data: data
      success: (data) ->
        if confirm "reload?"
          location.reload()
        else
          $("#myModal").modal "toggle"

  # delete
  $(".delete-comic").on "click", =>
    if not confirm("确定要删除?")
      false

    id = $(_this).data "id"
    $tr = $(_this).parent().parent()

    $.ajax
      type: "DELETE"
      url: "/api/comics/#{ id }"
      success: (data) ->
        if confirm("reload?")
          location.reload()
        else
          $tr.remove()
          $("#myModal").modal "toggle"
