const isOpenClass = "modal-is-open";
const openingClass = "modal-is-opening";
const closingClass = "modal-is-closing";
const scrollbarWidthCssVar = "--pico-scrollbar-width";
const animationDuration = 400; // ms
let visibleModal = null;

let copyToClipboard = (link) => {
    navigator.clipboard.writeText(link);
}

let setBatchSizeLabel = (unit) => {
    let input = document.querySelector('input[name="batch-size"]');
    let label = input.parentElement;

    label.childNodes[0].textContent = `Batch Size (${unit})`;
}

const setDescription = (event) => {
    let input = document.querySelector('textarea[name="description"]');
    input.value = event.target.innerText;
    closeModal(visibleModal);
}

const closeModal = (modal) => {
    visibleModal = null;
    const { documentElement: html } = document;
    html.classList.add(closingClass);
    setTimeout(() => {
      html.classList.remove(closingClass, isOpenClass);
      html.style.removeProperty(scrollbarWidthCssVar);
      modal.close();
    }, animationDuration);
  };

const openModal = (modal) => {
    const { documentElement: html } = document;
    const scrollbarWidth = getScrollbarWidth();
    if (scrollbarWidth) {
      html.style.setProperty(scrollbarWidthCssVar, `${scrollbarWidth}px`);
    }
    html.classList.add(isOpenClass, openingClass);
    setTimeout(() => {
      visibleModal = modal;
      html.classList.remove(openingClass);
    }, animationDuration);
    modal.showModal();
  };

  const toggleModal = (event) => {
    event.preventDefault();
    const modal = document.getElementById(event.currentTarget.dataset.target);
    if (!modal) return;
    modal && (modal.open ? closeModal(modal) : openModal(modal));
  };

  document.addEventListener("click", (event) => {
    if (visibleModal === null) return;
    const modalContent = visibleModal.querySelector("article");
    const isClickInside = modalContent.contains(event.target);
    !isClickInside && closeModal(visibleModal);
  });

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape" && visibleModal) {
      closeModal(visibleModal);
    }
  });
  
  const getScrollbarWidth = () => {
    const scrollbarWidth = window.innerWidth - document.documentElement.clientWidth;
    return scrollbarWidth;
  };
  
  const isScrollbarVisible = () => {
    return document.body.scrollHeight > screen.height;
  };