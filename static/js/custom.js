// OpenVPN UI helpers.
// AdminLTE 4 ships no jQuery, so everything here is plain DOM API.
// Bootstrap 5 components (Modal/Dropdown/Collapse/Tab) are driven by their
// data-bs-* attributes, and AdminLTE 4 widgets by data-lte-toggle — none of
// them need to be instantiated from JS here.

const MyAPP = {}

// Fire an API request the way the old $.ajax wrapper did: reload either way.
MyAPP.call = function (method, url, payload) {
  fetch(url, {
    method: method,
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload)
  })
    .then(function (response) { return response.json().catch(function () { return null }) })
    .then(function (data) { console.log(data) })
    .catch(function () { })
    .finally(function () { location.reload() })
}

MyAPP.Disconnect = function (cname) {
  console.log(cname)
  MyAPP.call("DELETE", "api/v1/session", { "cname": cname })
}

MyAPP.Restart = function (sname) {
  console.log(sname)
  MyAPP.call("POST", "api/v1/signal", { "sname": sname })
}

// Copy modal content to clipboard.
// custom.js is loaded with `defer`, so it runs after the document is parsed.
// The async Clipboard API is only exposed in secure contexts, but openvpn-ui is
// commonly served over plain HTTP on a LAN, so keep the legacy fallback.
function copyToClipboard(text) {
  if (navigator.clipboard && window.isSecureContext) {
    navigator.clipboard.writeText(text).catch(function () { })
    return
  }
  var area = document.createElement("textarea")
  area.value = text
  area.style.position = "absolute"
  area.style.left = "-9999px"
  document.body.appendChild(area)
  area.select()
  document.execCommand("copy")
  document.body.removeChild(area)
}

document.addEventListener("DOMContentLoaded", function () {
  document.querySelectorAll(".copy-modal-data-btn").forEach(function (button) {
    button.addEventListener("click", function () {
      var modal = button.closest(".modal")
      if (!modal) return
      var content = ""
      modal.querySelectorAll(".copy-details").forEach(function (el) {
        content += el.textContent.replace(/:\s/g, ": ") + "\n"
      })
      content = content.replace(/:\s\n/g, ": ")
      copyToClipboard(content)
    })
  })
})


function createEditor(name, size, theme, mode, readonly) {
  // find the textarea
  var textarea = document.querySelector("form textarea[name=" + name + "]");

  // create ace editor 
  var editor = ace.edit()
  editor.container.style.height = size

  editor.setTheme("ace/theme/" + theme); //"clouds_midnight"
  //editor.setTheme("ace/theme/twilight");
  //editor.setTheme("ace/theme/iplastic");

  editor.session.setMode("ace/mode/" + mode);

  editor.setReadOnly(readonly);
  editor.setShowPrintMargin(false);
  editor.session.setUseWrapMode(true);
  editor.session.setValue(textarea.value)
  // replace textarea with ace
  textarea.parentNode.insertBefore(editor.container, textarea)
  textarea.style.display = "none"
  // find the parent form and add submit event listener
  var form = textarea
  while (form && form.localName != "form") form = form.parentNode
  form.addEventListener("submit", function() {
      // update value of textarea to match value in ace
      textarea.value = editor.getValue()
  }, true)
}