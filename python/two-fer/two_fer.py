""" no parent module """

STRING_DEFAULT = "One for {}, one for me."


def two_fer(name="you"):
    """
    :param name: name to return as part of string
    """

    return STRING_DEFAULT.format(name)
