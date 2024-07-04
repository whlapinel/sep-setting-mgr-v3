# To Do

## PENDING

- dashboard: URL params should be changed to query params e.g. dashboard/classes/:class-id?students=true&test-events=true so that display state can be reflected accurately in the URL.
- test rooms functionality
- admin functionality

## COMPLETE

- change hx-post url in add-student-form component and url of handler
- add add student handler
- add students functionality: currently when you click add student, the form that is returned is the "add class" form. This needs to be changed to the add student form. This is because the Add Student button simply shows the hidden dialog with id of "dialog". So it opens the first dialog in the tree. Form component needs to be made so that more than one form can be present in the tree by modifying the Form (data) type and Form (templ) component, as well as the Javascript code in the Form (templ) component, so that clicking "Add Student" will show the dialog with a custom form id passed in  
- add signin check to signin page handler
