
# PENDING

# RESOLVED

- 7/11: canceling an edit form doesn't reset form, so changes persist on client side. Also true of add form. This will require the client to know the original values of those fields, which already sounds like it could get very complicated, and I'm trying to avoid lots of client-side code, I am wondering if I should change how the form is displayed. It might be faster and easier to fetch the form from the server with an hx-get rather than rendering all the forms in advance. When I tried it this way earlier I ran into problems with how the dialog is displayed, but perhaps I could send it to the same location and trigger an event on the client to open the dialog rather than trying to send it already open. So, create a custom event "showEditForm" on the client, a server-trigger via "HX-Trigger" that includes the form id e.g. "edit-student-form-12", and event listener that opens the closest dialog that is an ancestor of said form id.
- 7/2: unauthorized handler should be redirect so that hx-target doesn't render the page where it's expecting a component
- 5/16: users repo created twice (resolved 5/16)
