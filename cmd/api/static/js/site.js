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

document.addEventListener("DOMContentLoaded", (event) => {
  const srmCheckbox = document.querySelector("[name='useSrm']")
  const srmRange = document.querySelector("[name='srm']")

  if (srmCheckbox) {
    srmCheckbox.addEventListener('change', (event) => {
      if (event.currentTarget.checked) {
        srmRange.disabled = false;
      } else {
        srmRange.disabled = true;
      }
    })
  }
});


const getScrollbarWidth = () => {
  const scrollbarWidth = window.innerWidth - document.documentElement.clientWidth;
  return scrollbarWidth;
};

const isScrollbarVisible = () => {
  return document.body.scrollHeight > screen.height;
};

const srmMap = {};
srmMap["1"] = "#fee799"
srmMap["2"] = "#fdd978"
srmMap["3"] = "#fdcb5a"
srmMap["4"] = "#fcc143"
srmMap["5"] = "#f7b324"
srmMap["6"] = "#f5a800"
srmMap["7"] = "#ee9e00"
srmMap["8"] = "#e69100"
srmMap["9"] = "#e38901"
srmMap["10"] = "#da7e01"
srmMap["11"] = "#d37400"
srmMap["12"] = "#cb6c00"
srmMap["13"] = "#c66401"
srmMap["14"] = "#bf5c01"
srmMap["15"] = "#b65300"
srmMap["16"] = "#b04f00"
srmMap["17"] = "#ac4701"
srmMap["18"] = "#a24001"
srmMap["19"] = "#9c3900"
srmMap["20"] = "#963500"
srmMap["21"] = "#912f00"
srmMap["22"] = "#8c2d01"
srmMap["23"] = "#832501"
srmMap["24"] = "#7e1f01"
srmMap["25"] = "#771c01"
srmMap["26"] = "#721b00"
srmMap["27"] = "#6c1501"
srmMap["28"] = "#670f01"
srmMap["29"] = "#620f01"
srmMap["30"] = "#5b0d01"
srmMap["31"] = "#560c03"
srmMap["32"] = "#5d0a02"
srmMap["33"] = "#500a08"
srmMap["34"] = "#4a0605"
srmMap["35"] = "#440706"
srmMap["36"] = "#420807"
srmMap["37"] = "#3c0908"
srmMap["38"] = "#390708"
srmMap["39"] = "#39080b"
srmMap["40"] = "#35090a"