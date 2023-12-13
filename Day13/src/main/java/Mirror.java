import java.util.Arrays;
import java.util.Objects;

public class Mirror {
  public int[] position;
  public Orientation orientation;

  public Mirror(int[] position, Orientation orientation) {
    this.position = position;
    this.orientation = orientation;
  }

  int getSummary() {
    if (orientation == Orientation.VERTICAL) {
      return position[0];
    } else {
      return position[0]*100;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Mirror mirror = (Mirror) o;
    return Arrays.equals(position, mirror.position) && orientation == mirror.orientation;
  }

  @Override
  public int hashCode() {
    int result = Objects.hash(orientation);
    result = 31 * result + Arrays.hashCode(position);
    return result;
  }
}
