
$('.song').click(function(event) {
  event.preventDefault();
  $('source').remove();
  $('audio').append("<source src='" + $(this).attr('href') + "' </source>");
  //$('source').attr('src', $(this).attr('href')).trigger("load");
  $('source').trigger('load');
  $('audio').trigger('play');
});
