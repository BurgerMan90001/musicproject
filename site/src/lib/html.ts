export const escapeHTML = (v: string) => {
  var div = document.createElement("div");
  div.appendChild(document.createTextNode(v));
  return div.innerHTML;
};
