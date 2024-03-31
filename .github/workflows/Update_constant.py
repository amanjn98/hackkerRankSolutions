def update_golang_constant(file_path, constant_name, new_value):
    """
    Updates the value of a constant in a Golang file.

    Args:
        file_path: The path to the Golang file.
        constant_name: The name of the constant to update.
        new_value: The new value for the constant.
    """
    try:
        with open(file_path, "r") as file:
            lines = file.readlines()

        updated_lines = []
        for line in lines:
            if line.startswith(f"const {constant_name} "):
                # Update the line with the new value
                new_line = f"const {constant_name} = {new_value}\n"
            else:
                new_line = line

            updated_lines.append(new_line)

        with open(file_path, "w") as file:
            file.writelines(updated_lines)

        print(f"Successfully updated constant '{constant_name}' in '{file_path}' to {new_value}")
    except FileNotFoundError:
        print(f"Error: File not found: {file_path}")
    except Exception as e:
        print(f"Error updating constant: {e}")
