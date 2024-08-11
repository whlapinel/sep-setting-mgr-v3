(() => {
    const script = document.currentScript
    const select = script?.closest('select')
    select?.addEventListener('change', (event: any) => {
        console.log('dispatching optionSelected event')
        console.log('select', select)
        const selectedOption = select.options[event.target.selectedIndex]
        console.log('selectedOption', selectedOption)
        const roomSelectedEvent = new CustomEvent('optionSelected')
        selectedOption?.dispatchEvent(roomSelectedEvent)
    })
})()
