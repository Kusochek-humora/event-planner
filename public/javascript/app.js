'use strict'
document.addEventListener('DOMContentLoaded', function () {
    const answerWrapper = document.querySelectorAll('.answer-wrapper');
    const toggleBtns = document.querySelectorAll('.answer-toggle')

    for (const ans of answerWrapper) {
        ans.style.display = 'none';
    }

    for (const btn of toggleBtns) {
        btn.addEventListener('click', (e) => {
            const answer = e.target.parentElement.nextElementSibling;
            answer.style.display = answer.style.display === 'none' ? 'block' : 'none';
        });
    }


const editForm = document.querySelector('#form-update-task');
const taskToEdit = editForm && editForm.dataset.taskid;

if (editForm) {
    editForm.addEventListener('submit', async (event) => {
        event.preventDefault();

        const formData = Object.fromEntries(new FormData(editForm));

        try {
            const response = await fetch(`/task/${taskToEdit}`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(formData),
            });

            if (response.ok) {
                document.location.href = `/task/${taskToEdit}`;
            } else {
                console.error('Failed to update the task:', response.status);
            }
        } catch (error) {
            console.error('Error updating the task:', error.message);
        }
    });
}


const deleteButton = document.querySelector('#delete-button');
const taskToDelete = deleteButton && deleteButton.dataset.taskid;

if (deleteButton) {
    deleteButton.addEventListener('click', async () => {
        const result = confirm("Are you sure you want to delete this task?");

        if (!result) return;

        try {
            const response = await fetch(`/task/${taskToDelete}`, {
                method: 'DELETE'
            });

            if (response.ok) {
                document.location.href = "/";
            } else {
                console.error('Failed to delete the task:', response.status);
            }
        } catch (error) {
            console.error('Error deleting the task:', error.message);
        }
    });
}

})
