function Send(callback) {
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    callback();
  });
}

function Move() {
  location.href = './result.html?uuid=' + document.getElementById('uuid').value
}

var cntFormNum = 0;
function AddForm() {
  console.log(cntFormNum);
  var indentForm = '<div id = "languageIndnent' + String(cntFormNum) + '">'
    + '<p>言語(拡張子)</p><input type="text" id = "lang' + String(cntFormNum) + '"/>'
    + '<p>インデント数</p><input type="text" id = "indent' + String(cntFormNum) + '"/>'
    + '</div>';

  document.getElementById('languageIndent').insertAdjacentHTML('beforeend', indentForm);
  cntFormNum += 1;
  var element = document.getElementById("indentSettingNum");
  element.innerText = String(cntFormNum)
}
