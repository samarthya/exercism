""" no import used """


def response(hey_bob=None):
    """

    :param hey_bob: message that is conveyed to bob
    :returns string: response from bob
    """
    if hey_bob is not None:
        hey_bob = hey_bob.strip()

    if hey_bob is not None and len(hey_bob) == 0:
        return "Fine. Be that way!"

    if hey_bob.isupper() is False and hey_bob.endswith('?') is True:
        return 'Sure.'

    if hey_bob.isupper() is True and hey_bob.endswith('?') is True:
        return "Calm down, I know what I'm doing!"

    if hey_bob.isupper() is True:
        return 'Whoa, chill out!'

    return "Whatever."
