const rootId = "gofast-app";
const navigateHeader = "X-Gofast-Navigate";

function shouldHandleLink(event, anchor) {
  if (event.defaultPrevented || event.button !== 0) return false;
  if (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return false;
  if (anchor.target || anchor.hasAttribute("download")) return false;
  if (anchor.dataset.gofastIgnore !== undefined) return false;

  const url = new URL(anchor.href, window.location.href);
  return url.origin === window.location.origin && (url.pathname !== window.location.pathname || url.search !== window.location.search);
}

async function visit(url, options = {}) {
  const response = await fetch(url, {
    headers: { [navigateHeader]: "true" },
    credentials: "same-origin",
  });

  if (!response.ok) {
    window.location.href = url;
    return;
  }

  const html = await response.text();
  const root = document.getElementById(rootId);
  if (!root) {
    window.location.href = url;
    return;
  }

  root.innerHTML = html;
  const title = response.headers.get("X-Gofast-Title");
  if (title) document.title = title;
  if (!options.replace) history.pushState({}, "", url);
  window.scrollTo({ top: 0 });
}

document.addEventListener("click", (event) => {
  const anchor = event.target.closest("a[href]");
  if (!anchor || !shouldHandleLink(event, anchor)) return;
  event.preventDefault();
  visit(anchor.href).catch(() => {
    window.location.href = anchor.href;
  });
});

window.addEventListener("popstate", () => {
  visit(window.location.href, { replace: true }).catch(() => {
    window.location.reload();
  });
});
