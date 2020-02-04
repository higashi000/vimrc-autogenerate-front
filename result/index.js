var params = new URLSearchParams(window.location.search);
var vimrcUUID = params.get('uuid')

document.getElementById('vimrcUUID').value = vimrcUUID;

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
  go.run(result.instance);
});

function vimrcClipboard() {
  let vimrcEle = document.createElement("textarea");
  vimrcEle.value = document.getElementById("vimrc").innerText;

  document.body.appendChild(vimrcEle);
  vimrcEle.select();
  let res = document.execCommand("copy")

  vimrcEle.parentElement.removeChild(vimrcEle)

  if (res === true) {
    alert("クリップボードにコピーしました");
  } else {
    alert("クリップボードへのコピーに失敗しました");
  }
}
