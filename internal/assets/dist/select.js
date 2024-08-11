"use strict";
(() => {
    const script = document.currentScript;
    const select = script === null || script === void 0 ? void 0 : script.closest('select');
    select === null || select === void 0 ? void 0 : select.addEventListener('change', (event) => {
        console.log('dispatching optionSelected event');
        console.log('select', select);
        const selectedOption = select.options[event.target.selectedIndex];
        console.log('selectedOption', selectedOption);
        const roomSelectedEvent = new CustomEvent('optionSelected');
        selectedOption === null || selectedOption === void 0 ? void 0 : selectedOption.dispatchEvent(roomSelectedEvent);
    });
})();
