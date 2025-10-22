import logging
import os

import numpy as np
import tensorflow as tf
from tensorflow.python.client import device_lib


def printEnv():
    os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3'
    logging.getLogger('tensorflow').setLevel(logging.FATAL)
    gpus = tf.config.experimental.list_physical_devices('GPU')
    for gpu in gpus:
        tf.config.experimental.set_memory_growth(gpu, True)
    tf.debugging.set_log_device_placement(False)
    print(tf.__version__)
    np.set_printoptions(linewidth=200)
    print(device_lib.list_local_devices())