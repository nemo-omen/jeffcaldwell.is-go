class TestElement extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    const shadow = this.attachShadow({mode: "open"});
    const wrapper = document.createElement("div");
    wrapper.setAttribute("class", "whatever");

    const p = document.createElement("p");
    p.textContent = "And yet I live."
    wrapper.appendChild(p)

    shadow.appendChild(wrapper)
  }
}


customElements.define("test-element", TestElement);
