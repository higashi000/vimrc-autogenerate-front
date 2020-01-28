var params = new URLSearchParams(window.location.search);
var vimrcUUID = params.get('uuid')

document.getElementById('vimrcUUID').value = vimrcUUID;

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
  go.run(result.instance);
});
